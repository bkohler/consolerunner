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
			Pos:             Position{X: float64(m.rng.Intn(10)), Y: float64(initialY)}, // Cast ints to float64
			Velocity:        m.rng.Float64()*1.5 + 0.5,                                  // Random speed (0.5 to 2.0 cells/tick)
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
			// Cast termHeight and artHeight for comparison and assignment
			if m.runners[i].Pos.Y+float64(artHeight) >= float64(m.termHeight) {
				m.runners[i].Pos.Y = float64(m.termHeight - artHeight - 1)
				if m.runners[i].Pos.Y < 0 {
					m.runners[i].Pos.Y = 0.0 // Use float64 zero value
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
			// Position is already float64, just add velocity
			m.runners[i].Pos.X += m.runners[i].Velocity

			// Update animation frame
			m.runners[i].CurrentFrameIdx = (m.runners[i].CurrentFrameIdx + 1) % len(m.runners[i].ArtFrames)

			// Boundary check (wrap around screen width)
			artWidth := 0
			if len(m.runners[i].ArtFrames[m.runners[i].CurrentFrameIdx]) > 0 {
				artWidth = lipgloss.Width(m.runners[i].ArtFrames[m.runners[i].CurrentFrameIdx][0]) // Width of first line
			}
			// Cast termWidth and artWidth for comparison and assignment
			if m.runners[i].Pos.X > float64(m.termWidth) {
				m.runners[i].Pos.X = float64(-artWidth) // Reset position off-screen left
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

	// --- View implementation using a styled screen buffer ---

	// Helper struct to hold character and its style
	type StyledCell struct {
		Char  rune
		Style lipgloss.Style
	}

	// 1. Create the buffer (2D slice of StyledCell)
	defaultStyle := lipgloss.NewStyle() // Default style for empty cells
	buffer := make([][]StyledCell, m.termHeight)
	for y := 0; y < m.termHeight; y++ {
		buffer[y] = make([]StyledCell, m.termWidth)
		for x := 0; x < m.termWidth; x++ {
			buffer[y][x] = StyledCell{Char: ' ', Style: defaultStyle}
		}
	}

	// 2. Draw each runner onto the buffer, storing style
	for _, r := range m.runners {
		frame := r.ArtFrames[r.CurrentFrameIdx]
		// Determine the correct style based on theme
		runnerColor := r.Color.Light
		if lipgloss.HasDarkBackground() {
			runnerColor = r.Color.Dark
		}
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(runnerColor))

		for lineIdx, lineStr := range frame {
			// Cast lineIdx for calculation, cast result to int for buffer index
			targetY := int(r.Pos.Y + float64(lineIdx))
			if targetY < 0 || targetY >= m.termHeight {
				continue // Skip lines outside vertical bounds
			}

			currentXOffset := 0
			for _, char := range lineStr {
				charWidth := lipgloss.Width(string(char))
				// Cast currentXOffset for calculation, cast result to int for buffer index
				targetX := int(r.Pos.X + float64(currentXOffset))

				for i := 0; i < charWidth; i++ {
					// Cast i for calculation
					drawX := targetX + i // targetX is already int, i is int
					if drawX >= 0 && drawX < m.termWidth {
						cellChar := ' ' // Default for multi-width cells
						if i == 0 {
							cellChar = char
						}
						// Store character and its style
						// Ensure targetY is within buffer bounds (already checked)
						// Ensure drawX is within buffer bounds (already checked)
						buffer[targetY][drawX] = StyledCell{Char: cellChar, Style: style}
					}
				}
				// currentXOffset needs to track the float position for accurate placement
				// This buffer logic needs rethinking for float precision.
				// Let's revert targetX calculation and drawX calculation to use int(r.Pos.X) for now
				// and keep currentXOffset as int. This sacrifices sub-pixel accuracy in View
				// for simpler buffer indexing, while Update still uses floats.

				// Revert targetX calculation (line 198 change)
				// targetX := int(r.Pos.X) + currentXOffset

				// Revert drawX calculation (line 201 change) - it was already correct int+int
				// drawX := targetX + i

				// Keep currentXOffset as int
				currentXOffset += charWidth
			}
		}
	}

	// 3. Convert buffer to a single string, applying styles
	var finalView strings.Builder
	for y := 0; y < m.termHeight; y++ {
		for x := 0; x < m.termWidth; x++ {
			cell := buffer[y][x]
			finalView.WriteString(cell.Style.Render(string(cell.Char)))
		}
		// Add newline unless it's the last line
		if y < m.termHeight-1 {
			finalView.WriteString("\n")
		}
	}

	return finalView.String()

}

// tickCmd sends a tick message after a delay
func tickCmd() tea.Cmd {
	return tea.Tick(tickSpeed, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
