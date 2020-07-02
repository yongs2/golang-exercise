package main

import (
	"log"
)

func Lifo() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []string
	count int
}

func (s *Stack) Push(n string) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() string {
	if s.count == 0 {
		return ""
	}
	s.count--
	return s.nodes[s.count]
}

func main() {
	array := Lifo()
	array.Push("Amet")
	array.Push("ipsum")
	array.Push("Lorem")
	log.Printf("Pop=[%v], array=[%v][%v]\n", array.Pop(), array.count, array.nodes)
	log.Printf("Pop=[%v], array=[%v][%v]\n", array.Pop(), array.count, array.nodes)

	array.Push("sit")
	array.Push("Dolor")
	while := 1
	value := ""
	for while > 0 {
		value = array.Pop()
		if value == "" {
			while = 0
		} else {
			log.Printf("Pop=[%v], array=[%v][%v]\n", value, array.count, array.nodes)
		}
	}
	log.Printf("Done...\n")
}
