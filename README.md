# ai-commit

ai-commit is a command-line tool designed to help developers automatically generate meaningful commit messages for their Git repositories using AI. It uses the changes in the working directory to infer a commit message, promoting the use of conventional commits while keeping the messages precise and understandable.

![ai-commit](assets/ai-commit.gif)

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

Run the following command in your Git repository:

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

#### Docker

Run the next command to run over docker:

- Linux:
  ```sh
  docker run --network host -it --rm \
         -v ${HOME}/.config:/home/ai-commit/.config \
         -v $(pwd):/source \
         -e GIT_USER_NAME="$(git config --global user.name)" \
         -e GIT_USER_EMAIL="$(git config --global user.name)" \
         ai-commit:latest
  ```

- MacOS:
  ```sh
  docker run -it --rm \
         -v ${HOME}/.config:/home/ai-commit/.config \
         -v $(pwd):/source \
         -e GIT_USER_NAME="$(git config --global user.name)" \
         -e GIT_USER_EMAIL="$(git config --global user.name)" \
         ai-commit:latest --server http://host.docker.internal:11434
  ```

Two volumes are mounted:
- `${HOME}/.config:/home/ai-commit/.config`: To store the configuration file.
- `$(pwd):/source`: To access the source code.

Environment variables are set to configure the git user:
- `-e GIT_USER_NAME="$(git config --global user.name)"`: To set the git user name.
- `-e GIT_USER_EMAIL="$(git config --global user.email)"`: To set the git user email.

The `-it` flag is used to run the container in interactive mode to allow using `vim` to edit the commit message just 
generated. And the `--rm` flag is used to remove the container after it stops. `--network host` is used to allow the 
container to access the host network and connect to the Ollama server.

Any additional arguments passed to the `ai-commit` command will be forwarded to the Docker container. For example, to
run the command in noop mode `--noop` can be passed as an argument:

To make it easier you can create an alias in your `.bashrc` or `.zshrc` file:

```sh
alias ai-commit='docker run --network host -it --rm \
       -v ${HOME}/.config:/home/ai-commit/.config \
       -v $(pwd):/source \
       -e GIT_USER_NAME="$(git config --global user.name)" \
       -e GIT_USER_EMAIL="$(git config --global user.name)" \
       ai-commit:latest'
```

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

### Configuration

The configuration file `$(HOME)/.config/ai-commit/config.yaml` is automatically generated if it does not exist. This 
file contains the default values for the `ai-commit` configuration. Below is an example of the content of this file with 
the default values:

```yaml
logLevel: info
ollamaServer: http://localhost:11434
model: mistral
defaultPromptTemplate: |-
    Commit changes:
    {{ .Diff }}

    In an impersonal way, write a commit message that explains what the commit is for. Use conventional commits
    and the imperative mood in the first line. The first line should start with: feat, fix, refactor, docs, style,
    build, perf, ci, style, test or chore. Set the file name and the changes made in the body. Only one subject
    line is allowed. An example of commit message is:

    feat(file or class): Add user authentication

    - Implement user sign-up and login functionality
    - Add password hashing for security
    - Integrate with authentication API

    Add line breaks to separate subject from body.
defaultRetriesCommitMessage: 3
```

This file defines the following configuration parameters:

* `logLevel`: Logging level (default is "info").
* `ollamaServer`: URL of the Ollama server (default is http://localhost:11434).
* `model`: AI model to use for generating commit messages (default is "mistral").
* `defaultPromptTemplate`: Default template for the commit message.
* `defaultRetriesCommitMessage`: Number of retries for invalid commit messages (default is 3).
 
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

- `bin`: Compiles the `ai-commit` binary and places it in the `./bin` directory.
  ```sh
  make bin
  ```
- `install`: Builds the binary and copies it to `/usr/local/bin` for system-wide usage.
  ```sh
  make install
  ```
- `uninstall`: Removes the installed `ai-commit` binary from `/usr/local/bin`.
  ```sh
  make uninstall
- `release`: Make a release using goreleaser.
  ```sh
  make release
  ```
- `test-release`: Make a test release using goreleaser.
  ```sh
  make test-release
  ```
- `bump`: Bump the version of the project.
  ```sh
  make bump
  ```
- `clean`: Removes the `./bin` directory to clean up build artifacts.
  ```sh
  make clean
  ```

## Automatic Versioning

Just push a new tag with the version number and the CI/CD pipeline will take care of the rest.

Automatically bump the version with:
```sh
make new-version
```

If you want to manually bump the version:

```sh
cz bump --increment [MAJOR|MINOR|PATCH]
```

Then push the tag to the repository:

```sh
git push --tags
```

## License
[MIT](LICENSE.txt)

## Contributing
Feel free to submit pull requests or open issues to improve this project.
