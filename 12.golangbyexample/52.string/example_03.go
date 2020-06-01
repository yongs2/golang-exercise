package main

import (
	"log"
	"strconv"
)

func main() {
	example_ParseFloat()
	example_ParseBool()
}

func example_ParseFloat() {
	e1 := "1.3434"
	if s, err := strconv.ParseFloat(e1, 32); err == nil {
		log.Printf("32: %T, %v", s, s)
	}
	if s, err := strconv.ParseFloat(e1, 64); err == nil {
		log.Printf("64: %T, %v", s, s)
	}
}

func example_ParseBool() {
	var input string

	input = "true"
	if val, err := strconv.ParseBool(input); err == nil {
		log.Printf("[%v]: %T, %v", input, val, val)
	}

	input = "false"
	if val, err := strconv.ParseBool(input); err == nil {
		log.Printf("[%v]: %T, %v", input, val, val)
	}

	input = "garbage"
	if val, err := strconv.ParseBool(input); err == nil {
		log.Printf("[%v]: %T, %v", input, val, val)
	} else {
		log.Printf("[%v]: Given input is not a bool", input)
	}
}
