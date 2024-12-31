package cmd

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"

	"github.com/sergiotejon/ai-commit/internal/ai"
)

// checkCommitMessage checks the commit message for conventional commits.
func checkCommitMessage(message string) error {
	_, err := parser.NewMachine(parser.WithTypes(conventionalcommits.TypesConventional)).Parse([]byte(message))
	if err != nil {
		return err
	}

	return nil
}

// printCommitMessage prints the commit message.
func printCommitMessage(message string) {
	if !quiet && logLevel != "error" {
		fmt.Printf("\n%s\n\n", message)
	} else {
		fmt.Printf("%s\n", message)
	}
}

// generateCommitMessage generates a commit message using Ollama.
func generateCommitMessage(prompt, system, ollamaServer, model string, retries int, options ai.OllamaOptions) (string, error) {
	var result string
	var cleanResult string
	var err error

	for i := 0; i < retries; i++ {
		result, err = ai.QueryOllama(prompt, system, ollamaServer, model, options)
		if err != nil {
			// Retry if there is an error
			slog.Warn("Failed to query Ollama", "retry", i, "error", err)
			continue
		}

		cleanResult = strings.TrimSpace(strings.Map(func(r rune) rune {
			if strings.ContainsRune("`", r) {
				return -1
			}
			return r
		}, result))

		err = checkCommitMessage(cleanResult)
		if err == nil {
			// Valid commit message, exit the loop and continue
			break
		}
		printCommitMessage(cleanResult)
		slog.Error("Invalid commit message", "retry", i, "error", err)
	}

	if cleanResult == "" {
		return "", fmt.Errorf("failed to generate a valid commit message")
	}

	return cleanResult, err
}

// modifyCommitType modifies the commit type.
func modifyCommitType(commitMessage string, commitType string) string {
	parts := strings.SplitN(commitMessage, ":", 2)

	if len(parts) > 1 {
		scopeSplit := strings.SplitN(parts[0], "(", 2)
		if len(scopeSplit) > 1 {
			scope := strings.TrimRight(scopeSplit[1], ")")
			return fmt.Sprintf("%s(%s):%s", commitType, scope, parts[1])
		}
		return fmt.Sprintf("%s:%s", commitType, parts[1])
	}

	return commitMessage
}
