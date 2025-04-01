package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create the initial model
	m := initialModel()

	// Create and run the Bubble Tea program
	p := tea.NewProgram(m, tea.WithAltScreen()) // Use AltScreen for cleaner exit
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
