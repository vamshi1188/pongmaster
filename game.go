package main

import(
	"math/rand"
	"time"
)



// Game struct holds the main game state including the paddle, ball, score, and high score
type Game struct {
	paddle    Paddle
	ball      Ball
	score     int
	highScore int
}

// Layout defines the screen size, necessary for Ebiten to work properly
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}


// Update handles the game logic on each frame
func (g *Game) Update() error {	
	g.paddle.MoveOnKeyPress() // Move the paddle based on user input
	g.ball.Move()             // Move the ball according to its velocity
	g.CollideWithWall()       // Check for ball collisions with walls
	g.CollideWithPaddle()     // Check for ball collisions with the paddle
	return nil
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
