package main

import "fmt"
import "unicode/utf8"

func main() {
	var s1 string = "Hello world !!!"
	var s2 string = `Hello
	world !!!`
	var s3 string = "한글"

	fmt.Println("hello world !!!")
	fmt.Println(len(s1))	// 15
	fmt.Println(len(s2))	// 16
	fmt.Println(len(s3))	// 6 (unicode 이므로 한글 1자가 3bytes 로 처리됨)
	fmt.Println(utf8.RuneCountInString(s3))	// 2

	s1 = "Hello world ~ Wonderful"
	fmt.Println(s1)

	// String operation
	s1 = "Hello"
	s2 = "Hello"
	s3 = "World"
	fmt.Println(s1 == s2)
	fmt.Println(s1+s3)

	// boolean
	var V1 bool = true
	var V2 bool = false
	fmt.Println(V1, V2)

	// const
	const a1 = 10
	const b1 = 20
	fmt.Println(a1, b1)

	// 열거형
	const (
		a=iota	// 0
		b
		c
		d
		e
		f	
	)
	fmt.Println(a, b, c, d, e, f)
}
