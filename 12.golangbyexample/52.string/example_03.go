package main

import (
	"log"
	"strconv"
)

func main() {
	e1 := "1.3434"
	if s, err := strconv.ParseFloat(e1, 32); err == nil {
		log.Printf("32: %T, %v", s, s)
	}
	if s, err := strconv.ParseFloat(e1, 64); err == nil {
		log.Printf("64: %T, %v", s, s)
	}
}
