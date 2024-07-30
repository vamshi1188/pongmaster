package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"math/rand"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
	ballSpeed     = 4
	paddleSpeed   = 6
	highScoreFile = "highscore.txt"
)

type Object struct {
	X, Y, W, H int
}

type Paddle struct {
	Object
}

type Ball struct {
	Object
	dxdt int // x velocity per tick
	dydt int // y velocity per tick
}

type Game struct {
	paddle     Paddle
	ball       Ball
	score      int
	highScore  int
}

func main() {
	ebiten.SetWindowTitle("Pong in Ebitengine")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	paddle := Paddle{
		Object: Object{
			X: 600,
			Y: 200,
			W: 15,
			H: 100,
		},
	}

	ball := Ball{
		Object: Object{
			X: screenWidth / 2,
			Y: screenHeight / 2,
			W: 15,
			H: 15,
		},
	}
	rand.Seed(time.Now().UnixNano())
	ball.dxdt = (rand.Intn(2)*2 - 1) * ballSpeed // Randomize initial direction
	ball.dydt = (rand.Intn(2)*2 - 1) * ballSpeed

	g := &Game{
		paddle:    paddle,
		ball:      ball,
		highScore: loadHighScore(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

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

	// Draw the score and high score
	scoreStr := "Score: " + fmt.Sprint(g.score)
	text.Draw(screen, scoreStr, basicfont.Face7x13, 10, 10, color.White)

	highScoreStr := "High Score: " + fmt.Sprint(g.highScore)
	text.Draw(screen, highScoreStr, basicfont.Face7x13, 10, 30, color.White)
}

func (g *Game) Update() error {	
	g.paddle.MoveOnKeyPress() // Move the paddle based on user input
	g.ball.Move()             // Move the ball
	g.CollideWithWall()       // Check ball collisions with walls
	g.CollideWithPaddle()     // Check ball collisions with the paddle
	return nil
}

func (p *Paddle) MoveOnKeyPress() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.Y + p.H < screenHeight {
		p.Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.Y > 0 {
		p.Y -= paddleSpeed
	}
}

func (b *Ball) Move() {
	b.X += b.dxdt
	b.Y += b.dydt
}

func (g *Game) Reset() {
	g.ball.X = screenWidth / 2
	g.ball.Y = screenHeight / 2

	rand.Seed(time.Now().UnixNano())
	g.ball.dxdt = (rand.Intn(2)*2 - 1) * ballSpeed // Randomize reset direction
	g.ball.dydt = (rand.Intn(2)*2 - 1) * ballSpeed

	if g.score > g.highScore {
		g.highScore = g.score
		if err := saveHighScore(g.highScore); err != nil {
			log.Printf("Failed to save high score: %v", err)
		}
	}

	g.score = 0
}

func (g *Game) CollideWithWall() {
	if g.ball.X >= screenWidth {
		g.Reset()
	} else if g.ball.X <= 0 {
		g.ball.dxdt = ballSpeed
	} else if g.ball.Y <= 0 {
		g.ball.dydt = ballSpeed
	} else if g.ball.Y >= screenHeight {
		g.ball.dydt = -ballSpeed
	}
}

func (g *Game) CollideWithPaddle() {
	if g.ball.X >= g.paddle.X && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y + g.paddle.H {
		g.ball.dxdt = -g.ball.dxdt
		g.score++
		if g.score > g.highScore {
			g.highScore = g.score
		}
	}
}

// Save the high score to a file
func saveHighScore(score int) error {
	err := ioutil.WriteFile(highScoreFile, []byte(fmt.Sprint(score)), 0644)
	if err != nil {
		return fmt.Errorf("failed to save high score: %w", err)
	}
	return nil
}

// Load the high score from a file
func loadHighScore() int {
	data, err := ioutil.ReadFile(highScoreFile)
	if err != nil {
		if os.IsNotExist(err) {
			return 0 // No high score file found, return 0
		}
		log.Fatalf("Failed to load high score: %v", err)
	}
	var highScore int
	fmt.Sscanf(string(data), "%d", &highScore)
	return highScore
}
