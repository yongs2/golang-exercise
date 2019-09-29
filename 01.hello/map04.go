package main

import "fmt"

func main() {
	m := make(map[string] int)

	// 맵 m 의 요소를 삽입하거나 수정하기: m[key] = elem
	m["Answer"] = 42
	// 요소 값 가져오기: elem = m[key]
	fmt.Println("The Value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The Value:", m["Answer"])

	// 요소 지우기: delete(m, key)
	delete(m, "Answer")
	fmt.Println("The Value:", m["Answer"])

	// 키의 존재 여부 확인하기: elem, ok = m[key]
	v, ok := m["Answer"]
	fmt.Println("The Value:", v, "Present?", ok)
}