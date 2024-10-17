## 0.2.1 (2024-10-17)

### Fix

- **.github/workflows/release.yaml**: add GoReleaser execution

## 0.2.0 (2024-10-17)

### Feat

- **TODO.md**: Update task list and add completed CI/CD pipeline

## 0.1.3 (2024-10-17)

### Refactor

- **internal/git/commit.go**: Update git commit function to use fmt.Sprintf
- **internal/cmd/root.go**: Improve commit message generation

## 0.1.2 (2024-10-17)

### Fix

- **internal/cmd/root.go**: Modify commit handling
- **TODO.md**: Update TODO list

## 0.1.1 (2024-10-17)

### Fix

- **internal/git/diff.go**: retrieve root folder of repository

## 0.1.0 (2024-10-16)

### Feat

- **README.md**: Update README with installation instructions and dependencies

### Refactor

- **internal/cmd/root.go**: Modify command flags for better readability

## 0.0.1 (2024-10-16)

### Feat

- **internal/cmd/root.go**: Add version flag to ai-commit command for easy version checking
- **internal/version**: Create new file 'version.go'
- **internal/cmd/root.go**: Update project meta information.
- **internal/logger/logger.go**: Create new logger package for better application logging
- **internal/git/commit.go, internal/git/diff.go**: Introduce new files for commit and diff handling
- **internal/ai**: Create new file ollama.go
