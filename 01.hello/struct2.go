package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4	// 구조체에 속한 필드(데이터)는 dot(.) 으로 접근
	fmt.Println(v.X)
}