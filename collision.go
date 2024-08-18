package main

import (
	"math"

	
)

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

const maxReflectAngle = math.Pi / 4 // 45 degrees

// CollideWithPaddle handles ball collisions with the paddle
func (g *Game) CollideWithPaddle() {
	
		// Check for collision
		if g.ball.X >= g.paddle.X && g.ball.X <= g.paddle.X+g.paddle.W &&
		   g.ball.Y+g.ball.H >= g.paddle.Y && g.ball.Y <= g.paddle.Y+g.paddle.H {
		   
			// Calculate the hit position
			hitPosition := (float32(g.ball.Y) + float32(g.ball.H)/2) - (float32(g.paddle.Y) + float32(g.paddle.H)/2)
			hitFactor := hitPosition / float32(g.paddle.H/2)
			
			// Reflect the ball with an adjusted angle
			reflectAngle := hitFactor * maxReflectAngle // maxReflectAngle can be a predefined constant
			g.ball.dxdt = -g.ball.dxdt
			g.ball.dydt = int(float32(g.ball.dydt) + reflectAngle * float32(ballSpeed))
			

	
			g.score++ // Increase score
		}
	}
	
	

