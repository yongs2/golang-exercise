package main

import "fmt"

// struct 는 필드(데이터)들의 조합입니다.
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}