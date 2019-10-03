package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 메소드 리시버(method receiver) 는 func 키워드와 메소드의 이름 사이에 인자로 들어갑니다.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}

	fmt.Println(v.Abs())
}