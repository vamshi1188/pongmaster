package main


const (
	// Define screen dimensions and other constants
	screenWidth  = 640
	screenHeight = 480
	ballSpeed    = 4
	paddleSpeed  = 6
	highScoreFile = "highscore.txt" // File to save high score (not yet used)
)

// Object defines a general object with position and size
type Object struct {
	X, Y, W, H int
}
