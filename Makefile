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

.PHONY: all build run test fmt lint security-check clean install-tools

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