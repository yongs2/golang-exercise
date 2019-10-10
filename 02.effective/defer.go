package main

import "fmt"

func trace(s string) string {
	fmt.Println("Entering:", s)
	return s
}

func un(s string) {
	fmt.Println("Leaving:", s)
}

func a() {
	// defer 를 실행하는 함수가 반환되기 전에 즉각 함수 호출(연기된 함수)을 실행하도록 예약
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func multi_defer() {
	// 하나의 defer 호출 위치에서 여러개의 함수 호출을 지연할 수 있음
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func main() {
	b()
	multi_defer()
}
