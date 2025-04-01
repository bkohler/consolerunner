package main

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// RunnerType defines the different kinds of runners.
type RunnerType int

const (
	Jogger RunnerType = iota
	TrailRunner
	Marathoner
	CrewRunner
	UltraRunner
	TenKRunner
	// Add more types if needed
)

// Position represents the X, Y coordinates on the terminal screen.
type Position struct {
	X int
	Y int
}

// Runner represents a single animated runner on the screen.
type Runner struct {
	ID              int
	Type            RunnerType
	Pos             Position
	Velocity        float64    // Cells per tick (can be fractional for smoother perceived movement)
	ArtFrames       [][]string // Each inner slice is a frame, each string is a line of the frame
	CurrentFrameIdx int
	Color           lipgloss.AdaptiveColor // Use AdaptiveColor for better theme support
	rng             *rand.Rand             // Runner-specific RNG if needed, or use a global one
}

// Seed the global random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// String representation for RunnerType (optional but helpful for debugging)
func (rt RunnerType) String() string {
	switch rt {
	case Jogger:
		return "Jogger"
	case TrailRunner:
		return "TrailRunner"
	case Marathoner:
		return "Marathoner"
	case CrewRunner:
		return "CrewRunner"
	case UltraRunner:
		return "UltraRunner"
	case TenKRunner:
		return "TenKRunner"
	default:
		return "Unknown"
	}
}
