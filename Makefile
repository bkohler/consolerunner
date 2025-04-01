# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
BINARY_NAME=consolerunner
LINTCMD=golangci-lint
FMTCMD=goimports
GOSECCMD=gosec # Added for security checks

.PHONY: all build run test fmt lint security-check clean install-tools release

all: build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BINARY_NAME) .

# Run the application (builds first)
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Format code
fmt:
	@echo "Formatting code..."
	$(FMTCMD) -w .

# Run linter
lint:
	@echo "Running linter..."
	$(LINTCMD) run ./...

# Run security checks
security-check:
	@echo "Running security checks (gosec)..."
	$(GOSECCMD) ./...

# Clean build artifacts and logs
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME) narrator.log
	$(GOCLEAN)

# Install necessary Go tools (if not present)
install-tools:
	@echo "Installing tools (if needed)..."
	$(GOINSTALL) golang.org/x/tools/cmd/goimports@latest
	$(GOINSTALL) github.com/securego/gosec/v2/cmd/gosec@latest # Added gosec install
	# Assuming mise is used for golangci-lint as per previous steps
	# If not using mise, add manual install command here, e.g.:
	# $(GOINSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Ensure golangci-lint is installed (e.g., via mise use golangci-lint@latest or go install)"

# Create and publish a new release
release:
	@echo "Creating new release..."
	@LATEST_TAG=$$(gh release list --limit 1 --json tagName -q '.[0].tagName' || echo "v0.0.0")
	@echo "Latest tag: $${LATEST_TAG}"
	@# Increment version (simple increment of the middle number, assumes vX.Y.Z format)
	@NEXT_VERSION=$$(echo $${LATEST_TAG} | awk -F. '{printf "v%d.%d.%d", $$1, $$2+1, $$3}')
	@echo "Next version: $${NEXT_VERSION}"
	@# Define binary names
	@LINUX_BINARY=$(BINARY_NAME)-linux-amd64
	@WINDOWS_BINARY=$(BINARY_NAME)-windows-amd64.exe
	@CHECKSUM_FILE=checksums.txt
	@# Build binaries
	@echo "Building binaries..."
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $${LINUX_BINARY} .
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $${WINDOWS_BINARY} .
	@# Generate checksums
	@echo "Generating checksums..."
	sha256sum $${LINUX_BINARY} $${WINDOWS_BINARY} > $${CHECKSUM_FILE}
	@# Create GitHub release
	@echo "Creating GitHub release $${NEXT_VERSION}..."
	gh release create $${NEXT_VERSION} --generate-notes
	@# Upload assets
	@echo "Uploading assets..."
	gh release upload $${NEXT_VERSION} $${LINUX_BINARY} $${WINDOWS_BINARY} $${CHECKSUM_FILE}
	@# Clean up local files
	@echo "Cleaning up..."
	rm $${LINUX_BINARY} $${WINDOWS_BINARY} $${CHECKSUM_FILE}
	@echo "Release $${NEXT_VERSION} created successfully."
