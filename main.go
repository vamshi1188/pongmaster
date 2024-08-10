package main

import (
	"fmt"
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
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

// Ball represents the ball, using the Object struct and adding velocity fields
type Ball struct {
	Object
	dxdt int // x velocity per tick
	dydt int // y velocity per tick
}


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



// Draw renders the game objects on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the paddle
	vector.DrawFilledRect(screen, 
		float32(g.paddle.X), float32(g.paddle.Y), 
		float32(g.paddle.W), float32(g.paddle.H), 
		color.White, false,
	)

	// Draw the ball as a circle
	radius := float32(g.ball.W) / 2
	centerX := float32(g.ball.X) + radius
	centerY := float32(g.ball.Y) + radius
	vector.DrawFilledCircle(screen, centerX, centerY, radius, color.White, false)

	// Draw the score
	scoreStr := "Score: " + fmt.Sprint(g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 10, color.White)

	// Draw the high score
	highScoreStr := "High Score: " + fmt.Sprint(g.highScore)
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 30, color.White)
}



// Move updates the ball's position based on its velocity
func (b *Ball) Move() {
	b.X += b.dxdt
	b.Y += b.dydt
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

// CollideWithWall handles ball collisions with the screen boundaries
func (g *Game) CollideWithWall() {
	if g.ball.X >= screenWidth { // Ball hits the right wall
		g.Reset()
	} else if g.ball.X <= 0 { // Ball hits the left wall
		g.ball.dxdt = ballSpeed
	} else if g.ball.Y <= 0 { // Ball hits the top wall
		g.ball.dydt = ballSpeed
	} else if g.ball.Y >= screenHeight { // Ball hits the bottom wall
		g.ball.dydt = -ballSpeed
	}
}

// CollideWithPaddle handles ball collisions with the paddle
func (g *Game) CollideWithPaddle() {
	// Check if the ball is within the paddle's bounds
	if g.ball.X >= g.paddle.X && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y + g.paddle.H {
		g.ball.dxdt = -g.ball.dxdt // Reverse ball's horizontal direction
		g.score++ // Increase score
		if g.score > g.highScore { // Update high score if necessary
			g.highScore = g.score
		}
	}
}
