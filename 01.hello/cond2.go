package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {	// if 에서도 조건문 앞에 짧은 문장을 실행
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), 
		pow(3, 3, 20),
	)
}