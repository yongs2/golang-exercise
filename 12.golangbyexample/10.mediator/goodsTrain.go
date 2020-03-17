package main

import (
	"fmt"
)

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (g *goodsTrain) departure() {
	fmt.Println("GoodsTrain: Leaving")
	g.mediator.notifyFree()
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}
