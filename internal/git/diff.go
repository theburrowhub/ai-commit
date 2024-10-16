package git

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/utils/diff"
)

type FileDiff struct {
	File string
	Type string
	Text string
}

func GetDiffs() (string, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", fmt.Errorf("could not open the repository: %w", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		return "", fmt.Errorf("could not open the worktree: %w", err)
	}

	head, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("could not get the HEAD: %w", err)
	}

	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		return "", fmt.Errorf("could not get the commit object: %w", err)
	}

	tree, err := commit.Tree()
	if err != nil {
		return "", fmt.Errorf("could not get the tree: %w", err)
	}

	status, err := w.Status()
	if err != nil {
		return "", fmt.Errorf("could not get the status: %w", err)
	}

	// loop through the added and modified files
	var fileDiff []FileDiff
	for file, stat := range status {
		// TODO: Check delete
		if stat.Staging == git.Added || stat.Staging == git.Modified {
			// Retrieve the file from the index
			idxFile, err := w.Filesystem.Open(file)
			if err != nil {
				return "", fmt.Errorf("could not open the file %s: %w", file, err)
			}
			err = idxFile.Close()
			if err != nil {
				return "", fmt.Errorf("could not close the file %s: %w", file, err)
			}

			idxContent, err := os.ReadFile(file)
			if err != nil {
				return "", fmt.Errorf("could not read the file %s: %w", file, err)
			}

			// Retrieve the file from the last commit
			var commitContent string
			entry, err := tree.File(file)
			if err == nil {
				commitContent, err = entry.Contents()
				if err != nil {
					return "", fmt.Errorf("could not get the contents of the file %s: %w", file, err)
				}

				// Check the differences between the index and the last commit
				diffs := diff.Do(commitContent, string(idxContent))
				for _, d := range diffs {
					// Add the differences to the fileDiff slice
					fileDiff = append(fileDiff, FileDiff{
						File: file,
						Type: d.Type.String(),
						Text: d.Text,
					})
				}
			} else {
				fileDiff = append(fileDiff, FileDiff{
					File: file,
					Type: "NewFile",
					Text: "",
				})
			}
		} else if stat.Staging == git.Deleted {
			fileDiff = append(fileDiff, FileDiff{
				File: file,
				Type: "DeletedFile",
				Text: "",
			})
		}
	}

	// Convert the fileDiff slice to JSON
	diffsJSON, err := json.MarshalIndent(fileDiff, "", "  ")
	if err != nil {
		return "", fmt.Errorf("could not marshal the fileDiff: %w", err)
	}

	return string(diffsJSON), nil
}
