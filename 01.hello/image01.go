package main

import (
	"fmt"
	"image"
)

// Package image 는 Image 인터페이스를 정의합니다.
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}