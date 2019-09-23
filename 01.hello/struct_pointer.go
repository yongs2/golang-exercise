package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{ 1, 2 }
	
	q := &p		// Go에는 포인터가 있지만 포인터 연산은 불가능합니다.
			
	q.X = 1e9	// 구조체 변수는 구조체 포인터를 이용해서 접근 할 수 있습니다.
				// 포인터를 이용하는 간접적인 접근은 실제 구조체에도 영향을 미칩니다.

	fmt.Println(p)
}