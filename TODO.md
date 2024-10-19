# TODO

## Features

- [X] Configuration file for:
  - [X] AI model server URL
  - [X] Logging level
  - [X] Default AI model
- [X] Silent mode (no logs). Only errors to stderr. Useful along with `--noop` flag to work with automation.

## Improvements

- [X] Improve logging output
- [X] Improve AI prompt for better commit messages
- [X] Check commit messages if they adhere to conventional commit standards
  - [X] Retry if not adhering to standards (e.g., feat, fix, refactor)
  - [X] ... at least 3 times
- [X] Logs file (stderr and stdout). In case of error stderr.
- [ ] Tests

## Bug Fixes

- [X] Some commit messages are not being generated correctly 

## CI/CD

- [X] Add CI/CD pipeline to automate new releases
- [X] Docker image for easy installation
- [X] `make install` would install the binary and git extension with the current version
- [X] Commitizen support
