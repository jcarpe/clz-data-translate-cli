# Define variables
BINARY_NAME=clz-games-to-json-translator
GO_FILES=$(shell find . -name '*.go')

# Default target (execute tests)
all: test

# Format the code
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

# Run tests with coverage
test:
	@echo "Running tests..."
	@go test -v ./...

# Install dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod tidy

# Phony targets (not real files)
.PHONY: all fmt test deps
