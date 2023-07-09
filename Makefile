# Set up tools
.PHONY: install
install:
	sudo apt install pre-commit -y
	pre-commit --version
	pre-commit install
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
lint:
	golangci-lint run
