package main


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
