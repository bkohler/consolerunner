package main

// Placeholder ASCII art frames for different runner types.
// Each runner type has a slice of frames (for animation).
// Each frame is a slice of strings (representing lines of the art).
// IMPORTANT: Replace these placeholders with actual multi-line ASCII art!

var joggerArt = [][]string{
	{ // Frame 1
		" J ",
		"/|\\",
		"/ \\",
	},
	{ // Frame 2
		" J ",
		"\\|/",
		"\\ /",
	},
}

var trailRunnerArt = [][]string{
	{ // Frame 1
		" T ",
		"~|~",  // Wavy arms?
		"/^\\", // Rocky ground?
	},
	{ // Frame 2
		" T ",
		"~|~",
		"\\^/",
	},
}

var marathonerArt = [][]string{
	{ // Frame 1
		" M ",
		"-|-", // Steady arms
		"/ \\",
	},
	{ // Frame 2
		" M ",
		"-|-",
		"\\ /",
	},
}

var crewRunnerArt = [][]string{
	{ // Frame 1
		" C ",  // Cool C?
		"\\O/", // Arms up?
		" > >", // Fast legs?
	},
	{ // Frame 2
		" C ",
		"/O\\",
		"< <",
	},
}

var ultraRunnerArt = [][]string{
	{ // Frame 1
		" U ",
		"o|o", // Headlamp? Backpack?
		"_/\\_",
	},
	{ // Frame 2
		" U ",
		"o|o",
		"\\/_/",
	},
}

var tenKRunnerArt = [][]string{
	{ // Frame 1
		" 10K ", // Label?
		" /|>",
		" / >",
	},
	{ // Frame 2
		" 10K ",
		"<|\\ ",
		"< \\ ",
	},
}

// Map to easily access art based on RunnerType
var runnerArtMap = map[RunnerType][][]string{
	Jogger:      joggerArt,
	TrailRunner: trailRunnerArt,
	Marathoner:  marathonerArt,
	CrewRunner:  crewRunnerArt,
	UltraRunner: ultraRunnerArt,
	TenKRunner:  tenKRunnerArt,
}

// Function to get art for a runner type
func getArtForType(rt RunnerType) [][]string {
	art, ok := runnerArtMap[rt]
	if !ok {
		// Return a default or error indicator if type not found
		return [][]string{{"?"}}
	}
	return art
}
