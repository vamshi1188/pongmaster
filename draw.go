package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font/basicfont"
    "image/color"
    "fmt"
)

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