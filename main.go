package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenHeight = 1080
	screenWidth  = 1920
)

var (
	bgImage *ebiten.Image
)

type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage, nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}

	var err error
	bgImage, _, err = ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		log.Fatal(err)
	}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("meidän peli")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}

/*

	game := &Game{
		snake: NewSnake(),
		food:  NewFood(),
	}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}


*/
