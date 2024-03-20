package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Point struct {
	X int
	Y int
}

type Snake struct {
	Head      Point
	Direction Point
}

type Food struct {
	Position Point
}

type Game struct {
	snake         *Snake
	food          *Food
	score         int
	gameOver      bool
	updateCounter int
	speed         int
}

func NewSnake() *Snake {
	return &Snake{
		Head:      Point{X: screenWidth / tileSize / 2, Y: screenHeight / tileSize / 2},
		Direction: Point{X: 1, Y: 0},
	}
}

func NewFood() *Food {
	return &Food{
		Position: Point{
			X: rand.Intn(screenWidth / tileSize),
			Y: rand.Intn(screenHeight / tileSize),
		},
	}
}

func (s *Snake) Move() {
	s.Head.X += s.Direction.X
	s.Head.Y += s.Direction.Y
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}

	g.updateCounter++ /*
		if g.updateCounter < g.speed { // Why??????
			return nil
		}*/
	g.updateCounter = 0

	// Update the snake's position
	g.snake.Move()

	// Check if the snake hits the edges
	head := g.snake.Head
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth/tileSize || head.Y >= screenHeight/tileSize {
		g.gameOver = true
		g.speed = 10
	}

	// Calculate direction based on mouse movement
	mx, my := ebiten.CursorPosition()
	headX, headY := g.snake.Head.X*tileSize, g.snake.Head.Y*tileSize
	dx, dy := mx-headX, my-headY

	// Determine the direction of movement
	var direction Point
	if dx != 0 || dy != 0 {
		if dx > 0 {
			direction.X = 1
		} else if dx < 0 {
			direction.X = -1
		}
		if dy > 0 {
			direction.Y = 1
		} else if dy < 0 {
			direction.Y = -1
		}
	}

	// Update the snake's direction
	g.snake.Direction = direction

	//head := g.snake.Body[0]/*
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth/tileSize || head.Y >= screenHeight/tileSize {
		g.gameOver = false // if false, the snake will go over the edges
		g.speed = 20
	}

	// Check if the snake eats the food
	if head.X == g.food.Position.X && head.Y == g.food.Position.Y {
		g.score++ // Increment score once
		g.food = NewFood()
		if g.score == 20 { // when 20 points
			g.gameOver = true // You win!
		}
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw snake head
	head := g.snake.Head
	vector.DrawFilledRect(screen, float32(head.X*tileSize), float32(head.Y*tileSize), float32(tileSize), float32(tileSize), color.RGBA{255, 255, 0, 255}, true)

	// Draw food
	vector.DrawFilledRect(screen, float32(g.food.Position.X*tileSize), float32(g.food.Position.Y*tileSize), float32(tileSize), float32(tileSize), color.RGBA{255, 0, 0, 255}, true)

	// Create a font.Face
	face := basicfont.Face7x13

	// Draw game over text
	if g.gameOver {
		text.Draw(screen, "You win!", face, screenWidth/2-40, screenHeight/2, color.White)
		text.Draw(screen, "Press 'R' to restart", face, screenWidth/2-60, screenHeight/2+16, color.White)
	}

	// Draw score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, face, 5, screenHeight-5, color.White)

}
func (g *Game) restart() {
	g.snake = NewSnake()
	g.score = 0
	g.gameOver = false
	g.food = NewFood()
	g.speed = 20
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
