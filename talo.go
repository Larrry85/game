package main

import 

type Building struct {
	// Define building properties
	Position     Vector
	Health       int
	MaxHealth    int
	Production   string
	Construction int
}

type Vector struct {
	X, Y float64
}

var (
	buildings []*Building
	selected  *Building
)

