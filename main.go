package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenHeight = 1080
	screenWidth  = 1920
	tileSize     = 32
)

// BuildingType represents different types of buildings
type BuildingType int

const (
	// House represents a house building
	House BuildingType = iota
	// Factory represents a factory building
	Factory
)

// Building represents a building in the game world
type Building struct {
	Position     Vector
	Health       int
	MaxHealth    int
	Production   string
	Construction int
	Type         BuildingType // Type of building
}

// Vector represents a 2D vector
type Vector struct {
	X, Y float64
}

var (
	grassTile   *ebiten.Image
	buildingImg map[BuildingType]*ebiten.Image // Map to store building images
	buildings   []*Building
	selected    *Building
)

type Game struct{}

func (g *Game) Update() error {
	// Handle user input
	handleInput()

	// Update building logic
	for _, b := range buildings {
		if b.Production != "" {
			b.Construction += rand.Intn(5)
			if b.Construction >= 100 {
				b.Construction = 100
				b.Production = ""
			}
		}
	}

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

	rand.Seed(time.Now().UnixNano())

	// Create some buildings for demonstration
	buildings = append(buildings, &Building{Position: Vector{X: 100, Y: 100}, MaxHealth: 100, Health: 100, Type: House})
	buildings = append(buildings, &Building{Position: Vector{X: 200, Y: 200}, MaxHealth: 150, Health: 150, Production: "Unit", Type: Factory})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Demo Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
