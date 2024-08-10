package main

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
