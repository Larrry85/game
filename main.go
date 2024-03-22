package main

import (
	//"image/color" // only if there is background color
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// /////////
// Variables
// /////////
type BuildingType int

const ( // Houses
	House   BuildingType = iota
	Factory BuildingType = iota
)
const (
	screenHeight = 1080
	screenWidth  = 1920
	tileSize     = 32
)
const playerSpeed = 2 // Player's speed

var ( // for background and buildings
	grassTile   *ebiten.Image // background image type
	buildingImg map[BuildingType]*ebiten.Image
	buildings   []*Building
)
var ( // for Player
	playerImage *ebiten.Image // player image type
	isMoving    bool          // is the player moving?
	targetX     float64       // X-coordinate
	targetY     float64       // Y-coordinate
)

//////////
// Structs
//////////

// Building, types
type Building struct {
	Position     Vector
	Health       int
	MaxHealth    int
	Production   string
	Construction int
	Type         BuildingType
}

// coordinates on the 2D field
type Vector struct {
	X, Y float64
}

type Player struct {
	Image     *ebiten.Image // players type
	Position  Vector        // where player is, type
	Direction Vector        // where player moves, type
}

type Game struct {
	player    *Player       // player
	grassTile *ebiten.Image // background
	selected  *Building     // currently selected building
}

//////////
// Fuctions
///////////

func (g *Game) Update() error { // game keeps updating
	handleInput(g) // calls handleInput() to move the player
	// Update building logic
	for _, b := range buildings { // Dahl
		if b.Production != "" {
			b.Construction += rand.Intn(5)
			if b.Construction >= 100 {
				b.Construction = 100
				b.Production = ""
			} // Dahl END
		}
	}
	return nil
} //Update() END

func (g *Game) Draw(screen *ebiten.Image) { // uses Game struct
	// Clear the screen
	//screen.Fill(color.RGBA{R: 135, G: 206, B: 250, A: 255})
	// Draw background
	screen.DrawImage(grassTile, nil) // uses background image
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
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(b.Position.X, b.Position.Y)
		opts.GeoM.Scale(0.5, 0.5) // Scale the buildings
		screen.DrawImage(buildingImg[b.Type], opts)
	}
	/*
			//Dahl
		        // Draw construction progress bar
		        ebitenutil.DrawRect(screen, b.Position.X, b.Position.Y-5, float64(b.Construction)/100*64, 3, color.RGBA{0, 255, 0, 255})

		        // Draw selection box around selected building
		        if selected != nil && g.selected == b {
		            ebitenutil.DrawRect(screen, b.Position.X, b.Position.Y, 64, 64, color.RGBA{173, 216, 230, 255})
		        } //Dahl END*/
	// Draw player////////////////////////////////////////////
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.player.Position.X), float64(g.player.Position.Y))
	screen.DrawImage(playerImage, opts)
} // Draw() END

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
} // Layout() END

// Player moving
func handleInput(g *Game) { // uses Game struct
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) { // mouse click
		mouseX, mouseY := ebiten.CursorPosition()
		for _, b := range buildings { // Dahl
			if mouseX >= int(b.Position.X) && mouseX <= int(b.Position.X)+32 &&
				mouseY >= int(b.Position.Y) && mouseY <= int(b.Position.Y)+32 {
				g.selected = b
				break // Dahl END
			}
		}
		targetX = float64(mouseX)
		targetY = float64(mouseY)
		isMoving = true // Start moving towards mouse click
	}
	if isMoving { // boolean, player should stop when reaching the mouse click position
		// Calculate the direction towards the target position
		dx := targetX - g.player.Position.X
		dy := targetY - g.player.Position.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		// Set player's direction towards the target position
		if dist > 0 {
			g.player.Direction.X = dx / dist
			g.player.Direction.Y = dy / dist
		}
		// Move the player
		g.player.Position.X += g.player.Direction.X * playerSpeed
		g.player.Position.Y += g.player.Direction.Y * playerSpeed

		// Check if the player has reached the target position
		if dist <= playerSpeed {
			// Move the player directly to the target position
			g.player.Position.X = targetX
			g.player.Position.Y = targetY
			// Stop updating player's position
			isMoving = false // stop moving
		}
	}
	/* // Player can't go over the edges
	mouseX, mouseY := ebiten.CursorPosition()
	targetX := float64(mouseX)
	targetY := float64(mouseY)
	dx := targetX - g.player.Position.X
	dy := targetY - g.player.Position.Y
	dist := math.Sqrt(dx*dx + dy*dy)
	if dist <= playerSpeed {
		// Move the player directly to the target position
		g.player.Position.X = targetX
		g.player.Position.Y = targetY
		// Stop updating player's position
		g.player.Direction.X = 0
		g.player.Direction.Y = 0
	}*/
} // handleInput() END

func main() {
	player, _, err := ebitenutil.NewImageFromFile("player1.png")
	if err != nil {
		log.Fatal(err)
	}
	playerImage = player

	grass, _, err := ebitenutil.NewImageFromFile("background.png")
	if err != nil {
		log.Fatal(err)
	}
	grassTile = grass

	buildingImg = make(map[BuildingType]*ebiten.Image)
	house, _, err := ebitenutil.NewImageFromFile("talo.png")
	if err != nil {
		log.Fatal(err)
	}
	buildingImg[House] = house

	factory, _, err := ebitenutil.NewImageFromFile("talo.png")
	if err != nil {
		log.Fatal(err)
	}
	buildingImg[Factory] = factory

	rand.Seed(time.Now().UnixNano())

	buildings = append(buildings, &Building{Position: Vector{X: 100, Y: 100}, MaxHealth: 100, Health: 100, Type: House})
	buildings = append(buildings, &Building{Position: Vector{X: 200, Y: 200}, MaxHealth: 150, Health: 150, Type: Factory})

	game := &Game{
		player:    &Player{Image: playerImage},
		grassTile: grassTile,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Demo Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
} // main() END
