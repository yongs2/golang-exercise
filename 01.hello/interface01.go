package main

import (
	"fmt"
	"math"
)

// Abser 인터페이스는 메소드의 집합으로 정의됩니다.
type Abser interface {
	Abs() float64
}

// MyFloat 정의
type MyFloat float64

func main() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f	// a MyFloat implements Abser
	a = &v 	// a *Vertex implements Abser
	//a = v 	// a Vertex, cannot use v (type Vertex) as type Abser in assignment:

	fmt.Println(a.Abs())
	fmt.Println(f)
	fmt.Println(v)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}