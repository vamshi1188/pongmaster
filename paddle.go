package main

import "github.com/hajimehoshi/ebiten/v2"

type Paddle struct {
    Object
}

// MoveOnKeyPress moves the paddle based on arrow key input
func (p *Paddle) MoveOnKeyPress() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.Y + p.H < screenHeight {
		p.Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.Y > 0 {
		p.Y -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyU)&& p.Y > 0 {
		p.Y -= paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyI)&& p.Y + p.H < screenHeight {
		p.Y += paddleSpeed
	}
}