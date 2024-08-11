# Pong Game

This is a classic Pong game developed using [Ebitengine](https://ebitengine.org/), written in Go. The game includes a simple implementation of a single-player Pong game where the player controls a paddle to bounce the ball.

## Features

- **Single-player mode:** Control a paddle to keep the ball in play.
- **Score Tracking:** The game keeps track of the player's score.
- **High Score:** The game tracks and displays the highest score achieved during the session.
- **Dynamic Ball Movement:** The ball changes direction randomly on each reset.
- **Smooth Paddle Control:** Use arrow keys to move the paddle up and down.
- **Simple Graphics:** Rendered using Ebitengine with a customizable background and objects.

## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/vamshi1188/pongmaster.git
    cd pongmaster
    ```

2. **Install dependencies:**
    - Ensure you have Go installed. You can download it [here](https://golang.org/dl/).
    - Install the required Go modules:
      ```sh
      go get github.com/hajimehoshi/ebiten/v2
      go get golang.org/x/image/font/basicfont
      ```

3. **Run the game:**
    ```sh
    go run .
    ```

## Usage

- **Control the Paddle:** Use the **Up Arrow** and **Down Arrow** keys to move the paddle.
- **Reset the Game:** The game resets automatically when the ball passes the paddle.
- **High Score:** Try to beat your highest score!

## Code Structure

- **main.go:** Contains the main entry point and the game loop.
- **game.go:** Manages the game logic including ball movement, collision detection, and scoring.
- **paddle.go:** Handles paddle movement and input detection.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Future Enhancements (Optional Section)

- **Multi-player Mode:** Add a second paddle for competitive play.
- **Power-ups:** Introduce power-ups that affect the ball speed or paddle size.
- **Sound Effects:** Add sound effects for a more engaging experience.
- **Difficulty Levels:** Implement different difficulty levels with varying ball speeds.
