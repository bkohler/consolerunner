package main

import (
	"fmt"
	"time"
	"github.com/logrusorgru/aurora"
)

const (
	runnerFrame1 = "🏃"
	runnerFrame2 = "🏃‍♂️"
)

func main() {
	// Clear screen first
	fmt.Print("\033[H\033[2J")

	screenWidth := 50
	position := 0
	frame := 0
	au := aurora.NewAurora(true)

	// Animation loop
	for {
		// Clear line
		fmt.Print("\r")
		
		// Create empty space before runner
		for i := 0; i < position; i++ {
			fmt.Print(" ")
		}
		
		// Alternate between frames and colors
		if frame%2 == 0 {
			fmt.Print(au.Blue(runnerFrame1))
		} else {
			fmt.Print(au.Green(runnerFrame2))
		}
		
		// Update position and frame
		position = (position + 1) % screenWidth
		frame++
		
		// Control animation speed
		time.Sleep(100 * time.Millisecond)
	}
}