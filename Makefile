.PHONY: help bin install uninstall release test-release git-add-extension new-version clean

CURRENT_VERSION := $(shell git describe --tags --abbrev=0)

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Common targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

bin: ./bin/ai-commit ## Build go application

./bin/ai-commit:
	go build \
		-ldflags "-X github.com/theburrowhub/ai-commit/internal/version.version=${CURRENT_VERSION}" \
		-o ./bin/ai-commit ./cmd/main.go

install: bin /usr/local/bin/ai-commit /usr/local/bin/git-ai-commit ## Install ai-commit and git extension

/usr/local/bin/ai-commit:
	cp ./bin/ai-commit /usr/local/bin/ai-commit

/usr/local/bin/git-ai-commit:
	ln -s /usr/local/bin/ai-commit /usr/local/bin/git-ai-commit

uninstall:  ## Uninstall ai-commit and git extension
	rm /usr/local/bin/ai-commit
	rm /usr/local/bin/git-ai-commit

bump: ## Bump version using commitizen
	cz bump

bump-alpha: ## Bump version using commitizen (alpha)
	cz bump -pr alpha

bump-beta: ## Bump version using commitizen (beta)
	cz bump -pr beta

bump-rc: ## Bump version using commitizen (release candidate)
	cz bump -pr rc

release: ## Release new version
	goreleaser release

test-release: ## Test release new version
	goreleaser release --snapshot

clean: ## Clean up
	rm -rf ./bin
	rm -rf ./dist
