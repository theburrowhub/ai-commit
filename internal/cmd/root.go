package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/theburrowhub/ai-commit/internal/ai"
	"github.com/theburrowhub/ai-commit/internal/configure"
	"github.com/theburrowhub/ai-commit/internal/git"
	"github.com/theburrowhub/ai-commit/internal/logger"
	"github.com/theburrowhub/ai-commit/internal/version"
)

var (
	noop         bool
	logLevel     string
	ollamaServer string
	model        string
	retries      int
	showVersion  bool
	quiet        bool
	commitType   string
)

var rootCmd = &cobra.Command{
	Use:   "ai-commit",
	Short: "ai-commit is a tool to commit changes to a Git repository using AI for the commit message",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println("ai-commit version:", version.GetVersion())
			os.Exit(0)
		}

		// Setup the logger
		if quiet {
			logLevel = "error"
		}
		logger.SetupLogger(logLevel)

		// Get the differences in the repository
		diff, err := git.GetDiffs()
		if err != nil {
			logger.ErrorLogger.Error("Could not get the differences", "error", err)
			os.Exit(1)
		}

		// Check if there are changes to commit
		if diff == "null" {
			logger.ErrorLogger.Error("No changes to commit")
			os.Exit(1)
		}

		systemPrompt := string(*configure.Cfg.SystemPrompt)
		options := ai.OllamaOptions{
			NumCtx:      *configure.Cfg.NumCtx,
			Temperature: *configure.Cfg.Temperature,
			NumKeep:     *configure.Cfg.NumKeep,
		}

		slog.Debug("Prompt to infer", "systemPrompt", systemPrompt, "prompt", diff, "options", options)

		slog.Info("Querying Ollama...")
		slog.Info("Parameters", "server", ollamaServer, "model", model)

		// Generate the commit message using Ollama from the prompt with the differences
		commitMessage, err := generateCommitMessage(diff, systemPrompt, ollamaServer, model, retries, options)
		if err != nil {
			panic(err)
		}

		// If the commit type is not the default one, modify the commit type
		if commitType != "none" {
			commitMessage = modifyCommitType(commitMessage, commitType)
			slog.Debug("Modified commit message with new commit type")
		}

		slog.Debug("Commit message", "message", commitMessage)

		// If it's working in noop mode, print the commit message and exit
		if noop {
			slog.Warn("Running in noop mode, no changes made")
			printCommitMessage(commitMessage)
			os.Exit(0)
		}

		// If it's not in noop mode, commit the changes and amend the commit with the generated message
		err = git.Commit(commitMessage)
		if err != nil {
			panic(err)
		}
		err = git.CommitAmend()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	err := configure.LoadConfig()
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().BoolVar(&noop, "noop", *configure.Cfg.Noop, "Run without making any changes")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "logLevel", "l", *configure.Cfg.LogLevel, "Set the logging level (debug, info, warn, error)")
	rootCmd.PersistentFlags().StringVar(&ollamaServer, "server", *configure.Cfg.OllamaServer, "Set the Ollama server URL")
	rootCmd.PersistentFlags().StringVar(&model, "model", *configure.Cfg.Model, "Set the model to use for AI generation")
	rootCmd.PersistentFlags().IntVarP(&retries, "retries", "r", *configure.Cfg.RetriesCommitMessage, "Set the number of retries for invalid commit messages")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show the version of the application")
	rootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false, "Run in silent mode")
	rootCmd.PersistentFlags().StringVarP(&commitType, "commitType", "t", "none", "Set the type of the commit message")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
