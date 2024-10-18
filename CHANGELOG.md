## 0.5.7 (2024-10-18)

## 0.5.6 (2024-10-18)

## 0.5.5 (2024-10-18)

## 0.5.4 (2024-10-18)

## 0.5.3 (2024-10-18)

## 0.5.2 (2024-10-18)

### Refactor

- **.goreleaser.yaml**: Update Docker image repository URLs, modify release configurations

## 0.5.1 (2024-10-18)

### Fix

- **.goreleaser.yaml**: Update Docker image URLs and remove GitHub Container ...

## 0.5.0 (2024-10-18)

### Feat

- **TODO.md**: Update TODO list with completed and removed items

### Refactor

- **internal/cmd/root.go**: Modify error handling when retrieving Git repository diffs

## 0.4.0 (2024-10-17)

### Feat

- **internal/configure/configure.go**: Initialize new configure file
- **internal/configure/literalstring.go**: Introduce new file for literal string
- **internal/cmd/functions.go**: Implement conventional commits ...

### Fix

- **go.mod**: Update dependencies

### Refactor

- **internal/cmd/root.go**: Integrate configuration loading into ...

## 0.3.0 (2024-10-17)

## 0.2.2 (2024-10-17)

### Refactor

- **.github/workflows/release.yaml**: Update GoReleaser workflow

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
