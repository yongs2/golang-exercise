package main

import "fmt"

func main() {
	var z [] int

	// 슬라이스의 zero value는 nil 입니다.
	// nil 슬라이스는 길이와 최대 크기가 0입니다.
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil !!!")
	}
}