.PHONY: build

CURRENT_VERSION := $(shell git describe --tags --abbrev=0)

build:
	go build \
		-ldflags "-X github.com/sergiotejon/ai-commit/internal/version.version=${CURRENT_VERSION}" \
		-o ./bin/ai-commit ./cmd/main.go

install: build
	cp ./bin/ai-commit /usr/local/bin/ai-commit

uninstall:
	rm /usr/local/bin/ai-commit

release:
	goreleaser release

test-release:
	goreleaser release --snapshot

git-add-extension: install
	ln -s /usr/local/bin/ai-commit /usr/local/bin/git-ai-commit

new-version:
	cz bump

clean:
	rm -rf ./bin
	rm -rf ./dist