package main

import (
	"log"
	"math"
)

func main() {
	log.Printf("pi=[%v]", math.Pi)
	example_Ceil()
	example_Floor()
	example_Trunc()
	example_Round()
	example_RoundToEven()
	example_Abs()
	example_Sqrt()
	example_Cbrt()
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

func example_RoundToEven() {
	var f, res float64

	f = 0.5
	res = math.RoundToEven(f)
	log.Printf("RoundToEven(%v)=[%v]", f, res)

	f = 1.5
	res = math.RoundToEven(f)
	log.Printf("RoundToEven(%v)=[%v]", f, res)

	f = 2.5
	res = math.RoundToEven(f)
	log.Printf("RoundToEven(%v)=[%v]", f, res)

	f = 3.5
	res = math.RoundToEven(f)
	log.Printf("RoundToEven(%v)=[%v]", f, res)

	f = 4.5
	res = math.RoundToEven(f)
	log.Printf("RoundToEven(%v)=[%v]", f, res)
}

func example_Abs() {
	var f, res float64

	f = -2
	res = math.Abs(f)
	log.Printf("Abs(%v)=[%v]", f, res)

	f = 2
	res = math.Abs(f)
	log.Printf("Abs(%v)=[%v]", f, res)

	f = 3.5
	res = math.Abs(f)
	log.Printf("Abs(%v)=[%v]", f, res)

	f = -3.5
	res = math.Abs(f)
	log.Printf("Abs(%v)=[%v]", f, res)
}

func example_Sqrt() {
	var f, res float64

	f = 4
	res = math.Sqrt(f)
	log.Printf("Sqrt(%v)=[%v]", f, res)

	f = 9
	res = math.Sqrt(f)
	log.Printf("Sqrt(%v)=[%v]", f, res)

	f = 30.33
	res = math.Sqrt(f)
	log.Printf("Sqrt(%v)=[%v]", f, res)

	f = -9
	res = math.Sqrt(f)
	log.Printf("Sqrt(%v)=[%v]", f, res)
}

func example_Cbrt() {
	var f, res float64

	f = 8
	res = math.Cbrt(f)
	log.Printf("Cbrt(%v)=[%v]", f, res)

	f = 27
	res = math.Cbrt(f)
	log.Printf("Cbrt(%v)=[%v]", f, res)

	f = 30.33
	res = math.Cbrt(f)
	log.Printf("Cbrt(%v)=[%v]", f, res)
}
