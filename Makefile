.PHONY: build test clean run install build-all release help

# Build variables
BINARY_NAME=romkit
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_DIR=bin
MAIN_PATH=./cmd/cli/main.go

# Go build flags
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION)"

# Default target
all: build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME) $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for all platforms
build-all: clean
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)

	# Linux
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)

	# macOS
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)

	# Windows
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(MAIN_PATH)

	@echo "Build complete for all platforms"

# Run tests
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	@echo "Tests complete"

# Run tests with coverage report
test-coverage: test
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run the application
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Install the binary to GOPATH/bin
install:
	@echo "Installing $(BINARY_NAME)..."
	go install $(LDFLAGS) $(MAIN_PATH)
	@echo "Installation complete"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "Clean complete"

# Run go fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Format complete"

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...
	@echo "Vet complete"

# Run linters
lint: fmt vet
	@echo "Running linters..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it from https://golangci-lint.run/usage/install/"; \
	fi

# Create release archives
release: build-all
	@echo "Creating release archives..."
	@cd $(BUILD_DIR) && \
	for file in $(BINARY_NAME)-linux-* $(BINARY_NAME)-darwin-*; do \
		if [ -f "$$file" ]; then \
			tar -czf "$${file}.tar.gz" "$$file"; \
			echo "Created $${file}.tar.gz"; \
		fi; \
	done && \
	for file in $(BINARY_NAME)-windows-*.exe; do \
		if [ -f "$$file" ]; then \
			zip "$${file%.exe}.zip" "$$file"; \
			echo "Created $${file%.exe}.zip"; \
		fi; \
	done && \
	sha256sum *.tar.gz *.zip > checksums.txt 2>/dev/null || shasum -a 256 *.tar.gz *.zip > checksums.txt
	@echo "Release archives created in $(BUILD_DIR)/"

# Display help
help:
	@echo "RetroRomkit Makefile targets:"
	@echo ""
	@echo "  build         - Build binary for current platform"
	@echo "  build-all     - Build binaries for all platforms"
	@echo "  test          - Run tests with race detection"
	@echo "  test-coverage - Run tests and generate HTML coverage report"
	@echo "  run           - Build and run the application"
	@echo "  install       - Install binary to GOPATH/bin"
	@echo "  clean         - Remove build artifacts"
	@echo "  fmt           - Format Go code"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run all linters"
	@echo "  release       - Build binaries and create release archives"
	@echo "  help          - Display this help message"
	@echo ""
	@echo "Variables:"
	@echo "  VERSION       - Version to embed in binary (default: git describe)"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make test"
	@echo "  make VERSION=1.0.0 release"
