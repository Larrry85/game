/*package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

///////////
//Variables
///////////

//////////
// Structs
//////////

//////////
// Fuctions
///////////

func NewPlayer() *Player {
	playerImage, _, err := ebitenutil.NewImageFromFile("player.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Player{
		Image:    playerImage,
		Position: Vector{X: 100, Y: 100}, // Initial position of the player
	}
} // NewPlayer() END

func (p *Player) MoveTo(x, y int) {
	p.Position.X = float64(x)
	p.Position.Y = float64(y)
} // MoveTo() END

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.Image, op)
} // Draw() END*/
