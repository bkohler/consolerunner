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

*   Go (version 1.21 or later recommended)

### Installation & Running

1.  **Clone the repository (if you haven't already):**
    ```bash
    # git clone <repository-url>
    # cd consolerunner
    ```

2.  **Build the application:**
    ```bash
    go build
    ```

3.  **Run the executable:**
    ```bash
    ./consolerunner
    ```

4.  **Quit:** Press `q` or `Ctrl+C` to exit the application.

## ASCII Art

The ASCII art for the different runner types is defined in `art.go`. The current art is basic placeholder text. Feel free to replace it with your own creative, multi-frame ASCII animations!

## Development

*   **Dependencies:** Managed using Go modules (`go.mod`, `go.sum`).
*   **Linting:** Configured with `.golangci.yml`. Run `golangci-lint run ./...`
*   **Formatting:** Use `goimports` or `gofmt`.
*   **Testing:** Run tests with `go test ./...`