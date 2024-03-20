package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 10
	height = 10
)

type Entity struct {
	x, y int
}

type Game struct {
	player  Entity
	enemies []Entity
}

func (g *Game) Init() {
	g.player = Entity{x: width / 2, y: height / 2}
	g.generateEnemies(5)
}

func (g *Game) generateEnemies(numEnemies int) {
	for i := 0; i < numEnemies; i++ {
		rand.Seed(time.Now().UnixNano())
		enemy := Entity{x: rand.Intn(width), y: rand.Intn(height)}
		g.enemies = append(g.enemies, enemy)
	}
}

func (g *Game) Update() {
	// Update player position, for example, based on user input
	// For simplicity, we'll randomly move the player
	rand.Seed(time.Now().UnixNano())
	g.player.x += rand.Intn(3) - 1
	g.player.y += rand.Intn(3) - 1

	// Update enemies, for example, they can move towards the player
	for i := range g.enemies {
		if g.enemies[i].x < g.player.x {
			g.enemies[i].x++
		} else if g.enemies[i].x > g.player.x {
			g.enemies[i].x--
		}

		if g.enemies[i].y < g.player.y {
			g.enemies[i].y++
		} else if g.enemies[i].y > g.player.y {
			g.enemies[i].y--
		}
	}
}

func (g *Game) Render() {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// Render player
	grid[g.player.y][g.player.x] = 'P'

	// Render enemies
	for _, enemy := range g.enemies {
		grid[enemy.y][enemy.x] = 'E'
	}

	// Print the grid
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	game := Game{}
	game.Init()

	for {
		game.Update()
		game.Render()
		time.Sleep(500 * time.Millisecond) // Adjust for desired update rate
	}
}
