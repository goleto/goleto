.PHONY: fmt setup install-linter lint test build

setup: install-linter

install-linter:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: install-linter
	@golangci-lint run

test:
	@go test ./...

fmt:
	@go fmt ./...

build:
	@go build -o build/goleto ./cmd/goleto
