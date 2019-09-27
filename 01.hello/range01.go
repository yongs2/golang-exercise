package main

import "fmt"

var pow = [] int {1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// for 반복문에서 range 를 사용하면 슬라이스나 맵을 순회(iterates)할 수 있습니다.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}