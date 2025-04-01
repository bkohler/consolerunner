package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	minRunners = 3
	maxRunners = 8
	tickSpeed  = 100 * time.Millisecond // Animation frame rate
)

// Define messages
type tickMsg time.Time

// KeyMap for custom key bindings (optional but good practice)
type keyMap struct {
	Quit key.Binding
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
}

// model holds the application state
type model struct {
	runners    []Runner
	termWidth  int
	termHeight int
	keys       keyMap
	rng        *rand.Rand
	err        error // To store potential errors
}

// initialModel creates the starting state of the application
func initialModel() model {
	m := model{
		runners: make([]Runner, 0),
		keys:    keys,
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())), // Seed RNG
	}

	// Determine number of runners
	numRunners := m.rng.Intn(maxRunners-minRunners+1) + minRunners

	// Create runners
	for i := 0; i < numRunners; i++ {
		runnerType := RunnerType(m.rng.Intn(int(TenKRunner + 1))) // Random type up to TenKRunner
		art := getArtForType(runnerType)
		if len(art) == 0 || len(art[0]) == 0 {
			// Skip if art is invalid/empty
			continue
		}
		artHeight := len(art[0]) // Assuming all frames have same height

		// Random initial position (ensure within typical terminal height)
		// We'll adjust Y based on terminal height later in Update if needed
		initialY := m.rng.Intn(20) + 1 // Start between line 1 and 20 initially
		if initialY+artHeight > 24 {   // Avoid starting too low on common 24-line terms
			initialY = 24 - artHeight
		}
		if initialY < 0 {
			initialY = 0
		}

		newRunner := Runner{
			ID:              i,
			Type:            runnerType,
			Pos:             Position{X: m.rng.Intn(10), Y: initialY}, // Start near left edge
			Velocity:        m.rng.Float64()*1.5 + 0.5,                // Random speed (0.5 to 2.0 cells/tick)
			ArtFrames:       art,
			CurrentFrameIdx: 0,
			// Assign same random color for light/dark themes for simplicity
			Color: lipgloss.AdaptiveColor{Light: fmt.Sprintf("%d", m.rng.Intn(230)+16), Dark: fmt.Sprintf("%d", m.rng.Intn(230)+16)}, // Use color strings
			rng:   m.rng,                                                                                                             // Pass down RNG if needed per-runner
		}
		m.runners = append(m.runners, newRunner)
	}

	return m
}

// Init is the first command run by the Bubble Tea program.
func (m model) Init() tea.Cmd {
	return tickCmd() // Start the animation ticker
}

// Update handles messages and updates the model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		// Optional: Adjust runner Y positions if they are now off-screen due to resize
		for i := range m.runners {
			artHeight := len(m.runners[i].ArtFrames[0])
			if m.runners[i].Pos.Y+artHeight >= m.termHeight {
				m.runners[i].Pos.Y = m.termHeight - artHeight - 1
				if m.runners[i].Pos.Y < 0 {
					m.runners[i].Pos.Y = 0
				}
			}
		}
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

	case tickMsg:
		// Update runner positions and frames
		for i := range m.runners {
			// Calculate new position (simple linear movement for now)
			// Using float for position internally might be smoother, but int for rendering
			newX := float64(m.runners[i].Pos.X) + m.runners[i].Velocity
			m.runners[i].Pos.X = int(newX)

			// Update animation frame
			m.runners[i].CurrentFrameIdx = (m.runners[i].CurrentFrameIdx + 1) % len(m.runners[i].ArtFrames)

			// Boundary check (wrap around screen width)
			artWidth := 0
			if len(m.runners[i].ArtFrames[m.runners[i].CurrentFrameIdx]) > 0 {
				artWidth = lipgloss.Width(m.runners[i].ArtFrames[m.runners[i].CurrentFrameIdx][0]) // Width of first line
			}
			if m.runners[i].Pos.X > m.termWidth {
				m.runners[i].Pos.X = -artWidth // Reset position off-screen left
			}
		}
		return m, tickCmd() // Schedule next tick

	case error:
		m.err = msg
		return m, nil
	}

	return m, nil
}

// View renders the UI.
func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}
	if m.termWidth == 0 || m.termHeight == 0 {
		return "Initializing or terminal size too small..."
	}

	// --- Using a simpler View that just prints runners sequentially ---
	// This won't look like they are running across the screen correctly,
	// but avoids complex buffer logic for now. A proper implementation
	// would use a cell buffer or more advanced lipgloss techniques.
	var simpleView strings.Builder
	simpleView.WriteString(fmt.Sprintf("Console Runners! (%d runners) Press 'q' to quit.\n", len(m.runners)))
	simpleView.WriteString(fmt.Sprintf("Term size: %d x %d\n", m.termWidth, m.termHeight))
	simpleView.WriteString(strings.Repeat("-", m.termWidth) + "\n") // Separator line

	for _, r := range m.runners {
		frame := r.ArtFrames[r.CurrentFrameIdx]
		// Use the AdaptiveColor strings directly
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(r.Color.Light))
		if lipgloss.HasDarkBackground() {
			style = lipgloss.NewStyle().Foreground(lipgloss.Color(r.Color.Dark))
		}

		simpleView.WriteString(fmt.Sprintf("Runner %d (%s) at (%d, %d) V:%.1f\n", r.ID, r.Type, r.Pos.X, r.Pos.Y, r.Velocity))
		for _, line := range frame {
			// Basic rendering - does not handle positioning on screen
			simpleView.WriteString(style.Render(line) + "\n")
		}
		simpleView.WriteString(strings.Repeat("-", m.termWidth) + "\n") // Separator line
	}
	return simpleView.String() // Return the simple sequential view for now.

}

// tickCmd sends a tick message after a delay
func tickCmd() tea.Cmd {
	return tea.Tick(tickSpeed, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
