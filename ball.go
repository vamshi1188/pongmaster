package main

// Ball represents the ball, using the Object struct and adding velocity fields
type Ball struct {
	Object
	dxdt int // x velocity per tick
	dydt int // y velocity per tick
}

// Move updates the ball's position based on its velocity
func (b *Ball) Move() {
	b.X += b.dxdt
	b.Y += b.dydt
}