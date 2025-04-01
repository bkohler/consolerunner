# Running Man Animation in Go

This Go application creates a colorful animated running figure that moves across your terminal screen.

## Prerequisites

1. Make sure you have Go installed on your system:
   ```bash
   go version
   ```

2. The application requires a terminal that supports:
   - Unicode emoji characters
   - ANSI color codes
   - Cursor movement escape sequences

3. Required dependencies:
   - github.com/logrusorgru/aurora (for color support)

## Installation

1. Navigate to the project directory containing `main.go`

2. Install dependencies:
   ```bash
   go mod download
   ```

## Running the Application

You can run the application in one of two ways:

1. Compile and then execute:
   ```bash
   go build main.go
   ./main
   ```

2. Or compile and run in a single step:
   ```bash
   go run main.go
   ```

## What to Expect

When you run the program, you'll see:
- A running figure (🏃) moving across your terminal
- The figure alternates between blue and green colors
- The animation continues until you stop the program (Ctrl+C)

## Troubleshooting

If you see boxes or question marks instead of the runner emoji, your terminal might not support Unicode emoji characters.
If you don't see colors, your terminal might not support ANSI color codes or you might need to enable them.