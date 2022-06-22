package main

import "fmt"

func main() {
	game := newGame()

	// Add terrorists
	for i := 0; i < 4; i++ {
		game.addTerrorist(TerroristDressType)
	}

	// Add counter terrorists
	for i := 0; i < 4; i++ {
		game.addCounterTerrorist(CounterTerroristDressType)
	}

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\n DressColor: %s\n", dressType, dress.getColor())
	}
}
