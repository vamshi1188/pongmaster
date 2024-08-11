package main

import (
	"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Draw renders the game objects on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Set background color
	backgroundColor := color.RGBA{20, 10, 30,255} // Black background
	screen.Fill(backgroundColor) // Fill the entire screen with the background color

	// Draw the paddle
	paddleColor := color.RGBA{255, 255, 0, 255} // White color for paddle
	vector.DrawFilledRect(screen, 
		float32(g.paddle.X), float32(g.paddle.Y), 
		float32(g.paddle.W), float32(g.paddle.H), 
		paddleColor, false,
	)

	// Draw the ball as a circle
	ballColor := color.RGBA{0, 255, 0, 254} // Red color for ball
	radius := float32(g.ball.W) / 2
	centerX := float32(g.ball.X) + radius
	centerY := float32(g.ball.Y) + radius
	vector.DrawFilledCircle(screen, centerX, centerY, radius, ballColor, false)

	// Draw the score
	scoreStr := "Score: " + fmt.Sprint(g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 10, color.White)

	// Draw the high score
	highScoreStr := "High Score: " + fmt.Sprint(g.highScore)
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 30, color.White)
}
