package git

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/utils/diff"
)

type FileDiff struct {
	File string
	Type string
	Text string
}

// GetDiffs returns the differences between the files in the index and the last commit.
func GetDiffs() (string, error) {
	gitDir, err := findGitDir()
	if err != nil {
		return "", fmt.Errorf("could not find the .git directory: %w", err)
	}

	repo, err := git.PlainOpen(gitDir)
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

			f, err := os.Open(filepath.Join(gitDir, file))
			if err != nil {
				return "", fmt.Errorf("could not open the file %s: %w", file, err)
			}

			fileInfo, err := f.Stat()
			if err != nil {
				return "", fmt.Errorf("could not get the file info: %w", err)
			}

			var idxContent []byte
			if fileInfo.IsDir() {
				idxContent = []byte(fmt.Sprintf("New commits for submodule %s", file))
			} else {
				idxContent, err = os.ReadFile(filepath.Join(gitDir, file))
				if err != nil {
					return "", fmt.Errorf("could not read the file %s: %w", file, err)
				}
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
				fInfo, err := f.Stat()
				if err != nil {
					return "", fmt.Errorf("could not get the file info: %w", err)
				}
				if fInfo.IsDir() {
					fileDiff = append(fileDiff, FileDiff{
						File: file,
						Type: "Submodule",
						Text: fmt.Sprintf("Updated commits for submodule %s", file),
					})
				} else {
					fileDiff = append(fileDiff, FileDiff{
						File: file,
						Type: "NewFile",
						Text: "",
					})
				}
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

// findGitDir looks for the .git directory in the current directory or any of its parents.
func findGitDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get the current directory: %w", err)
	}

	for {
		gitPath := filepath.Join(dir, ".git")
		if _, err := os.Stat(gitPath); err == nil {
			pwd, err := os.Getwd()
			if err != nil {
				return "", fmt.Errorf("could not get the current directory: %w", err)
			}
			relPath, err := filepath.Rel(pwd, dir)
			if err != nil {
				return "", fmt.Errorf("could not get the relative path: %w", err)
			}
			return relPath, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("the .git directory was not found")
}
