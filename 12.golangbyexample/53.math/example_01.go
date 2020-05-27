package main

import (
	"log"
	"math"
)

func main() {
	example_Ceil()
	example_Floor()
	example_Trunc()
	example_Round()
}

// Ceil of a number is the least integer value greater than or equal to that number.
func example_Ceil() {
	var f, res float64

	f = 1.6
	res = math.Ceil(f)
	log.Printf("Ceil(%v)=[%v]", f, res)

	f = -1.6
	res = math.Ceil(f)
	log.Printf("Ceil(%v)=[%v]", f, res)

	f = 1
	res = math.Ceil(f)
	log.Printf("Ceil(%v)=[%v]", f, res)
}

// Floor of a number is the greatest integer value less than or equal to that number.
func example_Floor() {
	var f, res float64

	f = 1.6
	res = math.Floor(f)
	log.Printf("Floor(%v)=[%v]", f, res)

	f = -1.6
	res = math.Floor(f)
	log.Printf("Floor(%v)=[%v]", f, res)

	f = 1
	res = math.Floor(f)
	log.Printf("Floor(%v)=[%v]", f, res)
}

func example_Trunc() {
	var f, res float64

	f = 1.6
	res = math.Trunc(f)
	log.Printf("Trunc(%v)=[%v]", f, res)

	f = -1.6
	res = math.Trunc(f)
	log.Printf("Trunc(%v)=[%v]", f, res)

	f = 1
	res = math.Trunc(f)
	log.Printf("Trunc(%v)=[%v]", f, res)
}

func example_Round() {
	var f, res float64

	f = 1.6
	res = math.Round(f)
	log.Printf("Round(%v)=[%v]", f, res)

	f = 1.5
	res = math.Round(f)
	log.Printf("Round(%v)=[%v]", f, res)

	f = 1.4
	res = math.Round(f)
	log.Printf("Round(%v)=[%v]", f, res)

	f = -1.6
	res = math.Round(f)
	log.Printf("Round(%v)=[%v]", f, res)

	f = -1.5
	res = math.Round(f)
	log.Printf("Round(%v)=[%v]", f, res)

	f = -1.4
	res = math.Round(f)
	log.Printf("Round(%v)=[%v]", f, res)

	f = 1
	res = math.Round(f)
	log.Printf("Round(%v)=[%v]", f, res)
}
