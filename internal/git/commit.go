package git

import (
	"fmt"
	"os"
	"os/exec"
)

// Commit makes a new commit with the provided message
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not run git commit --amend: %w", err)
	}

	return nil
}

// CommitAmend runs git commit --amend
func CommitAmend() error {
	cmd := exec.Command("git", "commit", "--amend")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not run git commit --amend: %w", err)
	}

	return nil
}
