package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// 메소드는 구조체(struct) 뿐 아니라 아무 타입(type)에나 붙일 수 있습니다.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}