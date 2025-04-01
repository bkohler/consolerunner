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

// Position represents the X, Y coordinates on the terminal screen using float64 for smoother movement.
type Position struct {
	X float64
	Y float64
}

// Runner represents a single animated runner on the screen.
type Runner struct {
	ID              int
	Type            RunnerType
	Pos             Position
	VelocityX       float64    // Horizontal speed (cells per tick)
	VelocityY       float64    // Vertical speed (cells per tick)
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
