package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	data := make([][] uint8, dx)

	for x := range data {
		data[x] = make([] uint8, dy)
		for y := range data[x] {
			data[x][y] = uint8(x+y)	// 흥미로운 함수로는 x^y, (x+y)/2, x*y 등
		}
	}
	return data
}

func main() {
	pic.Show(Pic)
}