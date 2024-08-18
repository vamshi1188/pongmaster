package main

import (
	
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
	
	"math/rand"
	"time"
)



func main() {
	
	// Set the window title and size
	ebiten.SetWindowTitle("Pong in Ebitengine")
	ebiten.SetWindowSize(screenWidth, screenHeight)

	// Initialize the paddle
	paddle := Paddle{
		Object: Object{
			X: 700, // Paddle starts near the right edge
			Y: 200, // Initial vertical position
			W: 15,  // Width of the paddle
			H: 100, // Height of the paddle
		},
	}

	// Initialize the ball in the center of the screen
	ball := Ball{
		Object: Object{
			X: screenWidth /2, // Center of the screen
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












