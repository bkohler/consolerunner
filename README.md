# Console Runners

A fun Go application that displays various types of runners animating across your terminal screen using Bubble Tea and Lipgloss.

## Features

*   Displays multiple runners simultaneously.
*   Features different runner types (Joggers, Trail Runners, Marathoners, etc. - with placeholder ASCII art).
*   Randomized number, type, starting position, and speed for runners on each execution.
*   Uses Bubble Tea for the terminal UI framework.
*   Uses Lipgloss for styling.

## Getting Started

### Prerequisites

*   Go (version 1.21 or later recommended).
*   Optional: `mise` for managing Go and tool versions (as used during development).

### Installation & Running

1.  **Clone the repository (if you haven't already):**
    ```bash
    # git clone <repository-url>
    # cd consolerunner
    ```

2.  **Ensure Go tools are available:**
    *   If using `mise`, it might handle Go installation.
    *   Install `goimports` if needed: `go install golang.org/x/tools/cmd/goimports@latest`
    *   Install `golangci-lint` if needed (e.g., using `mise`): `mise use golangci-lint@latest`

3.  **Build the application (explicitly naming the output):**
    ```bash
    go build -o consolerunner .
    ```
    *(This ensures the executable is named `consolerunner`)*

4.  **Run the executable:**
    ```bash
    ./consolerunner
    ```

5.  **Quit:** Press `q` or `Ctrl+C` to exit the application.

## ASCII Art

The ASCII art for the different runner types is defined in `art.go`. The current art is basic placeholder text. Feel free to replace it with your own creative, multi-frame ASCII animations!

## Development

*   **Dependencies:** Managed using Go modules (`go.mod`, `go.sum`).
*   **Testing:** Run the unit tests using the standard Go command:
    ```bash
    go test ./...
    ```
*   **Formatting:** Use `goimports` (install if needed, see Prerequisites) to format code and manage imports:
    ```bash
    goimports -w .
    ```
    Alternatively, use the standard `gofmt`:
    ```bash
    gofmt -w .
    ```
*   **Linting:** Uses `golangci-lint` with its default settings (no `.golangci.yml` file is present). Install if needed (see Prerequisites) and run:
    ```bash
    golangci-lint run ./...