.PHONY: build

build:
	go build -o ./bin/ai-commit ./cmd/main.go

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

clean:
	rm -rf ./bin
	rm -rf ./dist