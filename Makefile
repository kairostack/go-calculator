.PHONY: all build test coverage clean install lint run help bench bench-operations fuzz fuzz-parser fuzz-operations

# Variables
BINARY_NAME=calculator
MAIN_PACKAGE=./cmd/calculator
BUILD_DIR=./build
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

# Default target
all: build

## Build the calculator binary
build:
	@echo "Building calculator..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Binary created at $(BUILD_DIR)/$(BINARY_NAME)"

## Install the calculator binary to GOPATH/bin
install: build
	@echo "Installing calculator..."
	go install $(MAIN_PACKAGE)
	@echo "Installed successfully"

## Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

## Run tests with race detection
test-race:
	@echo "Running tests with race detection..."
	go test -race -v ./...

## Run tests and generate coverage report
coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"
	@go tool cover -func=$(COVERAGE_FILE) | grep total

## Show coverage summary without generating HTML
coverage-summary:
	@echo "Running tests with coverage summary..."
	go test -cover ./...

## Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
	@echo "Clean complete"

## Run the calculator with example arguments
run: build
	@echo "Running calculator example:"
	$(BUILD_DIR)/$(BINARY_NAME) add 5 3

## Run the calculator help
run-help: build
	$(BUILD_DIR)/$(BINARY_NAME)

## Run linting (requires golangci-lint)
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		exit 1; \
	fi

## Format all Go code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

## Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

## Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

## Run the init command - initializes everything
init: deps build test
	@echo ""
	@echo "========================================="
	@echo "  Calculator Project Initialized!"
	@echo "========================================="
	@echo ""
	@echo "Available commands:"
	@echo "  make build          - Build the binary"
	@echo "  make test           - Run all tests"
	@echo "  make coverage       - Run tests with coverage"
	@echo "  make run            - Run example calculation"
	@echo "  make install        - Install to GOPATH/bin"
	@echo ""

## Display this help message
help:
	@echo "Available targets:"
	@grep -E '^## .*$$' $(MAKEFILE_LIST) | sed 's/## /  /'
	@echo ""
	@echo "Usage examples:"
	@echo "  make init              - Initialize project (deps, build, test)"
	@echo "  make build             - Build the calculator binary"
	@echo "  make test              - Run all unit tests"
	@echo "  make coverage          - Generate coverage report"
	@echo "  make run               - Run example: add 5 3"
	@echo "  ./build/calculator add 5 3  - Direct binary execution"
	@echo ""

# Development helpers
dev-add: build
	$(BUILD_DIR)/$(BINARY_NAME) add 10 5

dev-subtract: build
	$(BUILD_DIR)/$(BINARY_NAME) subtract 10 5

dev-multiply: build
	$(BUILD_DIR)/$(BINARY_NAME) multiply 10 5

dev-divide: build
	$(BUILD_DIR)/$(BINARY_NAME) divide 10 5

## Run all benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

## Run operation benchmarks only
bench-operations:
	@echo "Running operation benchmarks..."
	go test -bench=. -benchmem ./internal/operations/...

## Run all fuzz tests (30s each)
fuzz:
	@echo "Running fuzz tests (30s each)..."
	go test -fuzz=Fuzz -fuzztime=30s ./pkg/cli/... || true
	go test -fuzz=Fuzz -fuzztime=30s ./internal/operations/... || true

## Fuzz parser (60s)
fuzz-parser:
	@echo "Fuzzing parser (60s)..."
	go test -fuzz=FuzzParseArgs -fuzztime=60s ./pkg/cli/...

## Fuzz operations (60s)
fuzz-operations:
	@echo "Fuzzing operations (60s)..."
	go test -fuzz=Fuzz -fuzztime=60s ./internal/operations/...