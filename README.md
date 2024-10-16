# ai-commit

ai-commit is a command-line tool designed to help developers automatically generate meaningful commit messages for their Git repositories using AI. It uses the changes in the working directory to infer a commit message, promoting the use of conventional commits while keeping the messages precise and understandable.

## Features
- Generates commit messages using AI based on Git diffs.
- Adheres to conventional commit standards (e.g., feat, fix, refactor).
- Supports multiple logging levels (debug, info, warn, error).
- Option to run in "noop" mode to see the generated commit message without making changes.

## Requirements
- [Go](https://golang.org/) installed.
- [Cobra](https://github.com/spf13/cobra) library for the CLI.
- AI model server running at specified URL (Ollama server).

## Installation
Clone this repository and build the `ai-commit` binary:

```sh
git clone https://github.com/sergiotejon/ai-commit.git
cd ai-commit
make build
```

## Usage
The `ai-commit` command will use AI to generate a commit message based on the current state of your working directory.

### Basic Usage
```sh
ai-commit
```
This command will:
- Gather the current diffs in your working directory.
- Send the diffs to the configured AI model to generate a commit message.
- Commit the changes with the generated message.

### Command-Line Flags
- `--noop`: Run without making any changes to the Git repository (dry-run mode).
- `--logLevel`: Set the logging level (default is "info"). Available levels are "debug", "info", "warn", and "error".
- `--server`: Set the Ollama server URL (default is `http://localhost:11434`).
- `--model`: Specify the AI model to use for generating the commit message (default is "mistral").

### Example
```sh
ai-commit --noop --logLevel debug
```
This will run the command in noop mode, providing detailed logs about the process without actually committing anything.

## Commit Message Generation
The AI model uses the following prompt to generate commit messages:

```
In an impersonal way, write a commit message that explains what you did and why you did it.
Use conventional commits and the imperative mood in the first line.
The first line should start with: feat, fix, refactor, docs, style, build, perf, ci, style, test, or chore.
```

The generated commit message will adhere to the following format:
- A concise subject line following the conventional commit format.
- A detailed body explaining the changes made, if applicable.

### Dependencies
- AI communication and inference managed via internal package `ai`.
- Git operations performed using the internal package `git`.
- Logging provided via the internal package `logger`.

## Development
To execute the command during development:

```sh
go run ./cmd/ai-commit
```

### Makefile Targets
The Makefile provides the following targets to manage building, installing, and cleaning up:

- `build`: Compiles the `ai-commit` binary and places it in the `./bin` directory.
  ```sh
  make build
  ```
- `install`: Builds the binary and copies it to `/usr/local/bin` for system-wide usage.
  ```sh
  make install
  ```
- `uninstall`: Removes the installed `ai-commit` binary from `/usr/local/bin`.
  ```sh
  make uninstall
  ```
- `git-add-extension`: Installs `ai-commit` and creates a symlink to use it as a Git extension (`git ai-commit`).
  ```sh
  make git-add-extension
  ```
- `clean`: Removes the `./bin` directory to clean up build artifacts.
  ```sh
  make clean
  ```

## License
[MIT](LICENSE)

## Contributing
Feel free to submit pull requests or open issues to improve this project.
