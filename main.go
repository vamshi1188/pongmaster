package main

import (
	
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
	
	"math/rand"
	"time"
)

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

// Paddle represents the player's paddle, using the Object struct
// type Paddle struct {
// 	Object
// }



func main() {
	// Set the window title and size
	ebiten.SetWindowTitle("Pong in Ebitengine")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	// Initialize the paddle
	paddle := Paddle{
		Object: Object{
			X: 600, // Paddle starts near the right edge
			Y: 200, // Initial vertical position
			W: 15,  // Width of the paddle
			H: 100, // Height of the paddle
		},
	}

	// Initialize the ball in the center of the screen
	ball := Ball{
		Object: Object{
			X: screenWidth / 2, // Center of the screen
			Y: screenHeight / 2, // Center of the screen
			W: 15, // Width of the ball
			H: 15, // Height of the ball (ball is a square but drawn as a circle)
		},
	}
	// Randomize the initial direction of the ball
	rand.Seed(time.Now().UnixNano())
	ball.dxdt = (rand.Intn(2)*2 - 1) * ballSpeed // Randomly set dxdt to either -ballSpeed or ballSpeed
	ball.dydt = (rand.Intn(2)*2 - 1) * ballSpeed // Randomly set dydt to either -ballSpeed or ballSpeed

	// Create the game instance with the initialized paddle and ball
	g := &Game{
		paddle: paddle,
		ball:   ball,
	}

	// Start the game loop
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}









// Reset resets the ball's position and velocity, and updates the high score if needed
func (g *Game) Reset() {
	g.ball.X = screenWidth / 2
	g.ball.Y = screenHeight / 2

	// Randomize the direction of the ball after reset
	rand.Seed(time.Now().UnixNano())
	g.ball.dxdt = (rand.Intn(2)*2 - 1) * ballSpeed
	g.ball.dydt = (rand.Intn(2)*2 - 1) * ballSpeed

	// Update high score if the current score is higher
	if g.score > g.highScore {
		g.highScore = g.score
		// Save high score to a file (function not implemented yet)
		// saveHighScore(g.highScore)
	}

	g.score = 0 // Reset score to zero
}



