# ai-commit

ai-commit is a command-line tool designed to help developers automatically generate meaningful commit messages for their Git repositories using AI. It uses the changes in the working directory to infer a commit message, promoting the use of conventional commits while keeping the messages precise and understandable.

## Features
- Generates commit messages using AI based on Git diffs.
- Adheres to conventional commit standards (e.g., feat, fix, refactor).
- Supports multiple logging levels (debug, info, warn, error).
- Option to run in "noop" mode to see the generated commit message without making changes.

## Requirements
- AI model server running at specified URL ([Ollama](https://ollama.com) server).
- A AI Model to generate commit messages. By default, it uses the "mistral" model.
  ```sh
  ollama pull mistral
  ```
- Git installed and configured on the system.

## Installation

Run the following command to install the `ai-commit` binary and git extension:

```sh
curl -s https://raw.githubusercontent.com/sergiotejon/ai-commit/main/scripts/get-ai-commit.sh | sudo bash
```

### Clone and Build
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
or
```sh
git ai-commit
```

These commands will:
- Gather the current diffs in your working directory.
- Send the diffs to the configured AI model to generate a commit message.
- Commit the changes with the generated message.
- Open your default editor to allow you to modify the commit message before finalizing the commit.

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
go run ./cmd/main.go
```

### Requirements
- [Go](https://golang.org/) installed.
- [Cobra](https://github.com/spf13/cobra) library for the CLI.
- [go-git](https://github.com/go-git/go-git) library for interacting with Git repositories.

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
- `release`: Make a release using goreleaser.
  ```sh
  make release
  ```
- `test-release`: Make a test release using goreleaser.
  ```sh
  make test-release
  ```
- `clean`: Removes the `./bin` directory to clean up build artifacts.
  ```sh
  make clean
  ```

## License
[MIT](LICENSE.txt)

## Contributing
Feel free to submit pull requests or open issues to improve this project.
