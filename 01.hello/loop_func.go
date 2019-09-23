package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z = z - (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	var x float64 = 1

	for i := 0; i < 10; i++ {
		fmt.Printf("test.Sqrt(%g)=%g\n", x, Sqrt(x))
		fmt.Printf("math.Sqrt(%g)=%g\n", x, math.Sqrt(x))
		x += 1.0
	}
}