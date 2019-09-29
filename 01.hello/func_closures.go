package main

import "fmt"

// 함수는 클로져(full closures)
// 코드에서 adder 함수는 클로져(closure)를 반환
// 각각의 클로져는 자신만의 sum 변수를 가짐
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), 
			neg(-2*i),
		)
	}
}