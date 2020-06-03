package main

import (
	"log"
)

func main() {
	var shapes []shape

	s := &square{side: 2}
	shapes = append(shapes, s)

	r := &rectangle{length: 2, breath: 3}
	shapes = append(shapes, r)

	for _, shape := range shapes {
		log.Printf("Type: %s, Area %d\n", shape.getType(), shape.area())
	}
}

type shape interface {
	area() int
	getType() string
}

type rectangle struct {
	length int
	breath int
}

func (r *rectangle) area() int {
	return r.length * r.breath
}

func (r *rectangle) getType() string {
	return "rectangle"
}

type square struct {
	side int
}

func (s *square) area() int {
	return s.side * s.side
}

func (s *square) getType() string {
	return "square"
}
