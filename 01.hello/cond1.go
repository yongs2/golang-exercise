package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {	// if 문은 C와 Java와 비슷합니다. 조건 표현을 위해 ( ) 는 사용하지 않습니다. 하지만 실행문을 위한 { } 는 반드시 작성해야합니다.
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}