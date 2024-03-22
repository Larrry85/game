package main

import "math/rand"

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

var buildings []*Building

func createBuildings() {
	buildings = append(buildings, &Building{Position: Vector{X: 100, Y: 100}, MaxHealth: 100, Health: 100, Type: House})
	buildings = append(buildings, &Building{Position: Vector{X: 200, Y: 200}, MaxHealth: 150, Health: 150, Production: "Unit", Type: Factory})
}

func updateBuildings() {
	for _, b := range buildings {
		if b.Production != "" {
			b.Construction += rand.Intn(5)
			if b.Construction >= 100 {
				b.Construction = 100
				b.Production = ""
			}
		}
	}
}
