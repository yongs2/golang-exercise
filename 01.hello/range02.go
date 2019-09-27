package main

import "fmt"

func main() {
	pow := make([] int, 10)

	// 인덱스만 필요하다면 “, value” 부분을 다 제거
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	// _ 를 이용해서 인덱스(index)나 값(value)를 무시
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}