package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	tileSize     = 5
)

func main() {


	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebiten Web Browser Game")

	game := &Game{
		snake: NewSnake(),
		food:  NewFood(),
	}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}