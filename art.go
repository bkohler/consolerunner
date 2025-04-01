package main

// Placeholder ASCII art frames for different runner types.
// Each runner type has a slice of frames (for animation).
// Each frame is a slice of strings (representing lines of the art).
// IMPORTANT: Replace these placeholders with actual multi-line ASCII art!
// These are larger templates to guide detailed design.

var joggerArt = [][]string{
	{ // Frame 1
		"   ____  ",  // Hat/Hair
		"  / oo \\ ", // Face (eyes)
		" (|----|) ", // Face (mouth/jaw) / Torso
		" / |  | \\", // Arms / Torso
		" \\_|__|_/", // Legs / Torso
		"  /    \\ ", // Legs
		" /______\\", // Feet/Ground
		"         ",
	},
	{ // Frame 2
		"   ____  ",
		"  / oo \\ ",
		" (|----|) ",
		" \\ |  | /", // Arms moved
		" /_|__|_\\", // Legs moved
		"  \\    / ",
		" /______\\",
		"         ",
	},
	{ // Frame 3
		"   ____  ",
		"  / oo \\ ",
		" (|----|) ",
		" / |  | \\", // Arms back
		" \\_|__|_/", // Legs back
		"  /    \\ ",
		" /______\\",
		"         ",
	},
	{ // Frame 4
		"   ____  ",
		"  / oo \\ ",
		" (|----|) ",
		" \\ |  | /", // Arms moved again
		" /_|__|_\\", // Legs moved again
		"  \\    / ",
		" /______\\",
		"         ",
	},
}

// --- Define similar large, multi-frame placeholders for other runner types ---
// --- TrailRunner (e.g., add backpack, different terrain under feet) ---
var trailRunnerArt = joggerArt // Placeholder - Copy jogger for now

// --- Marathoner (e.g., add race bib, different posture) ---
var marathonerArt = joggerArt // Placeholder - Copy jogger for now

// --- CrewRunner (e.g., stylish clothes, headphones?) ---
var crewRunnerArt = joggerArt // Placeholder - Copy jogger for now

// --- UltraRunner (e.g., headlamp, hydration pack) ---
var ultraRunnerArt = joggerArt // Placeholder - Copy jogger for now

// --- TenKRunner (e.g., different build/speed posture) ---
var tenKRunnerArt = joggerArt // Placeholder - Copy jogger for now

// Map to easily access art based on RunnerType
var runnerArtMap = map[RunnerType][][]string{
	Jogger:      joggerArt,
	TrailRunner: trailRunnerArt, // Update with specific art
	Marathoner:  marathonerArt,  // Update with specific art
	CrewRunner:  crewRunnerArt,  // Update with specific art
	UltraRunner: ultraRunnerArt, // Update with specific art
	TenKRunner:  tenKRunnerArt,  // Update with specific art
}

// Function to get art for a runner type
func getArtForType(rt RunnerType) [][]string {
	art, ok := runnerArtMap[rt]
	if !ok || len(art) == 0 || len(art[0]) == 0 { // Check art validity
		// Return a default single-frame, single-line error indicator
		return [][]string{{"?"}}
	}
	return art
}
