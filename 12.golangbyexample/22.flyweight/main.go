package main

import (
	"fmt"
)

func main() {
	game := newGame()

	game.addTerrorlist(TerroristDressType)
	game.addTerrorlist(TerroristDressType)
	game.addTerrorlist(TerroristDressType)
	game.addTerrorlist(TerroristDressType)
	game.addTerrorlist(TerroristDressType)

	game.addCounterTerrorlist(CounterTerroristDressType)
	game.addCounterTerrorlist(CounterTerroristDressType)
	game.addCounterTerrorlist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}

	players := game.GetPlayers()
	for _, player := range players {
		fmt.Printf("player:%s, dress:%s\n", player.playerType, player.dress.getColor())
	}
}
