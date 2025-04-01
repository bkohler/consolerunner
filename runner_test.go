package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// Helper to create a basic runner for testing
func newTestRunner(x, y int, velocity float64, art [][]string) *Runner {
	// Ensure art is valid for testing frame logic
	if art == nil {
		art = [][]string{{"frame1"}, {"frame2"}} // Default 2-frame art
	}
	return &Runner{
		ID:              1,
		Type:            Jogger,                                 // Arbitrary type for testing
		Pos:             Position{X: float64(x), Y: float64(y)}, // Cast to float64
		VelocityX:       velocity,                               // Use VelocityX
		VelocityY:       0.0,                                    // No vertical movement in tests for simplicity
		ArtFrames:       art,
		CurrentFrameIdx: 0,
		Color:           lipgloss.AdaptiveColor{Light: "1", Dark: "1"},
		rng:             rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func TestUpdatePosition(t *testing.T) {
	tests := []struct {
		name         string
		runner       *Runner
		termWidth    int
		expectedPosX float64 // Change expected type to float64
	}{
		{
			name:         "Move right within bounds",
			runner:       newTestRunner(10, 5, 2.0, [][]string{{"art"}}),
			termWidth:    80,
			expectedPosX: 12.0, // 10 + 2.0
		},
		{
			name:         "Move right with fractional velocity",
			runner:       newTestRunner(10, 5, 1.5, [][]string{{"art"}}),
			termWidth:    80,
			expectedPosX: 11.5, // 10 + 1.5
		},
		{
			name:         "Wrap around screen edge",
			runner:       newTestRunner(78, 5, 3.0, [][]string{{"abc"}}), // Art width 3
			termWidth:    80,
			expectedPosX: -3.0, // 78 + 3 = 81 > 80, wraps to float64(-artWidth)
		},
		{
			name:         "Start off screen left, move right",
			runner:       newTestRunner(-5, 5, 2.0, [][]string{{"art"}}),
			termWidth:    80,
			expectedPosX: -3.0, // -5 + 2.0
		},
		{
			name:         "Nil runner",
			runner:       nil,
			termWidth:    80,
			expectedPosX: 0.0, // Use float64 zero value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Need to handle the nil case explicitly before calling updatePosition
			if tt.runner == nil {
				updatePosition(tt.runner, tt.termWidth) // Call with nil to check for panic
				// No position check needed for nil runner
			} else {
				originalY := tt.runner.Pos.Y // updatePosition shouldn't change Y
				updatePosition(tt.runner, tt.termWidth)
				// Compare float64 values. For simple cases, direct comparison is often okay.
				// For more complex calculations, consider using a tolerance (e.g., math.Abs(a-b) < epsilon).
				if tt.runner.Pos.X != tt.expectedPosX {
					t.Errorf("updatePosition() Pos.X = %v, want %v", tt.runner.Pos.X, tt.expectedPosX)
				}
				if tt.runner.Pos.Y != originalY {
					t.Errorf("updatePosition() changed Pos.Y unexpectedly to %v", tt.runner.Pos.Y)
				}
			}
		})
	}
}

func TestNextFrame(t *testing.T) {
	twoFrameArt := [][]string{{"f1"}, {"f2"}}
	threeFrameArt := [][]string{{"f1"}, {"f2"}, {"f3"}}

	tests := []struct {
		name              string
		runner            *Runner
		expectedFrameIdx  int
		expectedFrameIdx2 int // After second call
	}{
		{
			name:              "Cycle through 2 frames",
			runner:            newTestRunner(0, 0, 1.0, twoFrameArt),
			expectedFrameIdx:  1, // 0 -> 1
			expectedFrameIdx2: 0, // 1 -> 0
		},
		{
			name:              "Cycle through 3 frames",
			runner:            newTestRunner(0, 0, 1.0, threeFrameArt),
			expectedFrameIdx:  1, // 0 -> 1
			expectedFrameIdx2: 2, // 1 -> 2
		},
		{
			name:              "Single frame art",
			runner:            newTestRunner(0, 0, 1.0, [][]string{{"f1"}}),
			expectedFrameIdx:  0, // 0 -> 0
			expectedFrameIdx2: 0, // 0 -> 0
		},
		{
			name:              "Nil runner",
			runner:            nil,
			expectedFrameIdx:  0, // Expect no change/panic
			expectedFrameIdx2: 0,
		},
		{
			name:              "Empty art frames",
			runner:            newTestRunner(0, 0, 1.0, [][]string{}),
			expectedFrameIdx:  0, // Expect no change/panic
			expectedFrameIdx2: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.runner == nil {
				nextFrame(tt.runner) // Call 1
				nextFrame(tt.runner) // Call 2
				// No index check needed
			} else {
				originalIndex := tt.runner.CurrentFrameIdx
				nextFrame(tt.runner) // Call 1
				if tt.runner.CurrentFrameIdx != tt.expectedFrameIdx {
					t.Errorf("nextFrame() call 1 index = %v, want %v", tt.runner.CurrentFrameIdx, tt.expectedFrameIdx)
				}
				nextFrame(tt.runner) // Call 2
				if tt.runner.CurrentFrameIdx != tt.expectedFrameIdx2 {
					t.Errorf("nextFrame() call 2 index = %v, want %v", tt.runner.CurrentFrameIdx, tt.expectedFrameIdx2)
				}
				// Reset for next test iteration if needed, though runner is created fresh each time
				tt.runner.CurrentFrameIdx = originalIndex
			}
		})
	}
}
