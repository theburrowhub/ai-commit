package cmd

import (
	"fmt"
	"github.com/sergiotejon/ai-commit/internal/version"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sergiotejon/ai-commit/internal/ai"
	"github.com/sergiotejon/ai-commit/internal/git"
	"github.com/sergiotejon/ai-commit/internal/logger"
)

// TODO: configure
const (
	prompt = "" +
		"In an impersonal way, write a commit message that explains what the commit is for. " +
		"Use conventional commits and the imperative mood in the first line. " +
		"The first line should start with: feat, fix, refactor, docs, style, build, perf, ci, style, test or chore. " +
		"Set the file name and the changes made in the body. " +
		"Only one subject line is allowed. " +
		"An example of commit message is: " +
		"" +
		"feat(file or class): Add user authentication\\n" +
		"\\n" +
		"- Implement user sign-up and login functionality\\n" +
		"- Add password hashing for security\\n" +
		"- Integrate with authentication API\\n" +
		"- ...\\n" +
		"" +
		"Add line breaks to separate subject from body."
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

		diff, err := git.GetDiffs()
		if err != nil {
			panic(err)
		}

		logger.SetupLogger(logLevel)

		slog.Debug("Prompt to infer", "prompt", prompt, "diff", diff)

		if diff == "null" {
			slog.Info("No changes to commit")
			os.Exit(0)
		}

		slog.Info("Querying Ollama...")
		slog.Info("Parameters", "server", ollamaServer, "model", model)
		result, err := ai.QueryOllama(fmt.Sprintf("Commit changes:\n%s\n%s", diff, prompt), ollamaServer, model)
		if err != nil {
			panic(err)
		}

		cleanResult := strings.TrimSpace(strings.Map(func(r rune) rune {
			if strings.ContainsRune("`", r) {
				return -1
			}
			return r
		}, result))

		slog.Debug("Commit message", "message", cleanResult)

		if !noop {
			err = git.Commit(cleanResult)
			if err != nil {
				panic(err)
			}

			err = git.CommitAmend()
			if err != nil {
				panic(err)
			}
		} else {
			slog.Info("Running in noop mode, no changes made")
			fmt.Println("\nCommit message:\n\n", cleanResult)
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
