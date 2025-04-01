package main

import "github.com/charmbracelet/lipgloss"

// updatePosition calculates the new position of a runner based on its velocity.
// For now, it only updates the X coordinate. Y coordinate updates could be added later.
// Note: This function currently modifies the runner directly. A pure function returning
// the new position might be preferable for testability.
func updatePosition(runner *Runner, termWidth int) {
	if runner == nil {
		return
	}

	// Calculate new X position based on velocity
	newX := float64(runner.Pos.X) + runner.Velocity
	runner.Pos.X = int(newX)

	// Boundary check (wrap around screen width)
	artWidth := 0
	if len(runner.ArtFrames) > 0 && runner.CurrentFrameIdx < len(runner.ArtFrames) && len(runner.ArtFrames[runner.CurrentFrameIdx]) > 0 {
		// Use lipgloss.Width for accurate display width calculation
		artWidth = lipgloss.Width(runner.ArtFrames[runner.CurrentFrameIdx][0]) // Width of first line
	} else {
		// Default width if art is invalid or empty
		artWidth = 1
	}

	if runner.Pos.X > termWidth {
		runner.Pos.X = -artWidth // Reset position off-screen left
	}
}

// nextFrame calculates the next animation frame index for a runner.
// Note: This function currently modifies the runner directly.
func nextFrame(runner *Runner) {
	if runner == nil || len(runner.ArtFrames) == 0 {
		return
	}
	runner.CurrentFrameIdx = (runner.CurrentFrameIdx + 1) % len(runner.ArtFrames)
}

// --- Helper functions for runner creation (can be expanded) ---

// getRandomRunnerAttributes generates random attributes for a new runner.
// This could be moved from model.go or enhanced here.
// func getRandomRunnerAttributes(rng *rand.Rand) (RunnerType, Position, float64, lipgloss.AdaptiveColor) {
// 	// ... implementation ...
// }
