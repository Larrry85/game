package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenHeight = 1080
	screenWidth  = 1920
	tileSize     = 32
)

var (
	grassTile   *ebiten.Image
	buildingImg map[BuildingType]*ebiten.Image // Map to store building images
	selected    *Building
)

type Game struct{}

func (g *Game) Update() error {
	// Handle user input
	handleInput()

	// Update building logic
	updateBuildings()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{R: 135, G: 206, B: 250, A: 255})

	// Draw map tiles
	op := &ebiten.DrawImageOptions{}
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			op.GeoM.Reset()
			op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
			screen.DrawImage(grassTile, op)
		}
	}

	// Draw buildings
	for _, b := range buildings {
		// Draw building image at its position
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(b.Position.X, b.Position.Y)
		screen.DrawImage(buildingImg[b.Type], opts)

		// Draw construction progress bar
		ebitenutil.DrawRect(screen, b.Position.X, b.Position.Y-5, float64(b.Construction)/100*64, 3, color.RGBA{0, 255, 0, 255})

		// Draw selection box around selected building
		if selected != nil && selected == b {
			ebitenutil.DrawRect(screen, b.Position.X, b.Position.Y, 64, 64, color.RGBA{173, 216, 230, 255})
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func handleInput() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		for _, b := range buildings {
			if mouseX >= int(b.Position.X) && mouseX <= int(b.Position.X)+32 &&
				mouseY >= int(b.Position.Y) && mouseY <= int(b.Position.Y)+32 {
				selected = b
				break
			}
		}
	}
}

func main() {
	game := &Game{}

	var err error

	grassTile, _, err = ebitenutil.NewImageFromFile("grass.png")
	if err != nil {
		log.Fatal(err)
	}

	// Load building images
	buildingImg = make(map[BuildingType]*ebiten.Image)
	buildingImg[House], _, err = ebitenutil.NewImageFromFile("talo.png")
	if err != nil {
		log.Fatal(err)
	}
	buildingImg[Factory], _, err = ebitenutil.NewImageFromFile("talo.png")
	if err != nil {
		log.Fatal(err)
	}

	createBuildings()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Demo Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
