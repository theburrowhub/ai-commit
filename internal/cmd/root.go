package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/sergiotejon/ai-commit/internal/git"
	"github.com/sergiotejon/ai-commit/internal/logger"
	"github.com/sergiotejon/ai-commit/internal/version"
)

// TODO: configure
const (
	retriesCommitMessage = 3

	promptTemplate = `
        Commit changes:
		{{ .Diff }}

		In an impersonal way, write a commit message that explains what the commit is for. Use conventional commits
        and the imperative mood in the first line. The first line should start with: feat, fix, refactor, docs, style,
        build, perf, ci, style, test or chore. Set the file name and the changes made in the body. Only one subject
		line is allowed. An example of commit message is:

		feat(file or class): Add user authentication

		- Implement user sign-up and login functionality
		- Add password hashing for security
		- Integrate with authentication API

		Add line breaks to separate subject from body.
    `
)

var (
	noop         bool
	logLevel     string
	ollamaServer string
	model        string
	showVersion  bool
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
		logger.SetupLogger(logLevel)
		// Load configuration
		// ...

		// Get the differences in the repository
		diff, err := git.GetDiffs()
		if err != nil {
			panic(err)
		}

		// Check if there are changes to commit
		if diff == "null" {
			slog.Info("No changes to commit")
			os.Exit(0)
		}

		// Return prompt to infer
		prompt, err := generatePrompt(promptTemplate, diff)
		if err != nil {
			panic(err)
		}

		slog.Debug("Prompt to infer", "prompt", prompt)

		slog.Info("Querying Ollama...")
		slog.Info("Parameters", "server", ollamaServer, "model", model)

		// Generate the commit message using Ollama from the prompt with the differences
		commitMessage, err := generateCommitMessage(prompt, ollamaServer, model, retriesCommitMessage)
		if err != nil {
			panic(err)
		}

		slog.Debug("Commit message", "message", commitMessage)

		// If it's working in noop mode, print the commit message and exit
		if noop {
			slog.Info("Running in noop mode, no changes made")
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
	rootCmd.PersistentFlags().BoolVar(&noop, "noop", false, "Run without making any changes")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "logLevel", "l", "info", "Set the logging level (debug, info, warn, error)")
	rootCmd.PersistentFlags().StringVar(&ollamaServer, "server", "http://localhost:11434", "Set the Ollama server URL")
	rootCmd.PersistentFlags().StringVar(&model, "model", "mistral", "Set the model to use for AI generation")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show the version of the application")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
