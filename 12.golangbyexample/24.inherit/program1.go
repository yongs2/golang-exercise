package main

import (
	"fmt"
)

type iBase interface {
	say()
}

type base struct {
	color string
}

func (b *base) say() {
	b.clear()
}

func (b *base) clear() {
	fmt.Println("Clear from base's function")
}

type child struct {
	base
	style string
}

func (b *child) clear() {
    fmt.Println("Clear from child's function")
}

func check(b iBase) {
	b.say()
}

func main() {
	base := base{color: "Red"}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
	fmt.Println("The color is " + child.color)
}
