package main

import (
	"fmt"
)

type iAlpha interface {
	work()
	common(iAlpha)
}

type alpha struct {
	name string
}

func (a *alpha) common(i iAlpha) {
	fmt.Println("a.common called, name=", a.name)
	i.work()
}

type beta struct {
	alpha
}

func (b *beta) work() {
	fmt.Println("b.work called")
	fmt.Printf("b.name is %s\n", b.name)
}

func main() {
	a := alpha{
		name: "test",
	}
	b := &beta{
		alpha: a,
	}
	b.common(b)
}
