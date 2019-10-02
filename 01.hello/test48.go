package main

import (
	"fmt"
	"math/cmplx"
)

// https://go-tour-kr.appspot.com/#48
// 복소수 세제곱근
// 세제곱근을 얻기 위해서는, 뉴턴의 방법 (Newton's method)을 적용하여 다음을 반복 수행합니다:
// z = z - (z * z * z - x) / (3 * z * z)
func Cbrt(x complex128) complex128 {
	var z complex128 = 1.

	for i := 0; i < 10; i++ {
		z = z - (z * z * z - x) / (3 * z * z)
	}

	fmt.Println("Pow=", cmplx.Pow(z, 3))
	return z
}

func main() {
	// 2의 세제곱근
	fmt.Println(Cbrt(2))
}