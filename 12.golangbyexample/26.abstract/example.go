package main

import (
	"fmt"
)

type iAlpha interface {
	work()
	common()
}

type alpha struct {
	name string
}

func (a *alpha) common() {
	fmt.Println("common called")
}

type beta struct {
	alpha
}

func (b *beta) work() {
	fmt.Println("work called")
	fmt.Printf("name is %s\n", b.name)
	b.common()
}

func main() {
	a := alpha{
		name: "test",
	}
	b := &beta{
		alpha: a,
	}
	b.work()
}
