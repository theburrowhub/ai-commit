package cmd

import (
	"bytes"
	"fmt"
	"log/slog"
	"strings"
	"text/template"

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
	fmt.Printf("\nCommit message:\n\n%s\n\n", message)
}

// generatePrompt generates a prompt template using the diff.
func generatePrompt(promptTemplate, diff string) (string, error) {
	var tmpl, err = template.New("prompt").Parse(promptTemplate)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, struct{ Diff string }{Diff: diff})
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}

// generateCommitMessage generates a commit message using Ollama.
func generateCommitMessage(prompt, ollamaServer, model string, retries int) (string, error) {
	var result string
	var cleanResult string
	var err error

	for i := 0; i < retries; i++ {
		result, err = ai.QueryOllama(prompt, ollamaServer, model)
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
