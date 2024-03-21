package main

import "errors"

type Building interface {
	Build() error
	Upgrade() error
	Destroy() error
	// Other methods as needed
}

func handleBuildingInteraction(x, y int) {
	// Check if the click is within the bounds of any building
	for _, building := range buildings {
		if building.Contains(x, y) {
			// Handle building interaction
			// For example: building.Select(), building.Upgrade(), etc.
			return
		}
	}
}

func (b *Base) Upgrade() error {
	// Check if player has enough resources
	if player.Resources >= b.UpgradeCost {
		// Deduct resources
		player.Resources -= b.UpgradeCost
		// Upgrade building
		b.Level++
		return nil
	}
	return errors.New("insufficient resources")
}

func canPlaceBuilding(x, y int) bool {
	// Check if the position is within the map bounds
	if x < 0 || y < 0 || x >= MapWidth || y >= MapHeight {
		return false
	}
	// Check if there are any obstacles or other buildings at the position
	for _, obstacle := range obstacles {
		if obstacle.Contains(x, y) {
			return false
		}
	}
	// Check other conditions as needed
	return true
}
