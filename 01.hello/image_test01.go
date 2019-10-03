package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

// 당신 자신의 Image 타입을 정의하시고, 필수 함수들 을 구현하신 다음, pic.ShowImage 를 호출하십시오.
type Image struct {
	width int
	height int
}

// Bounds 는 image.Rect(0, 0, w, h) 와 같은 image.Rectangle 을 반환해야 합니다.
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// ColorModel 은 color.RGBAModel 을 반환해야 합니다.
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// At 은 하나의 컬러를 반환해야 합니다; 지난 그림 생성기에서 값 v 는 color.RGBA{v, v, 255, 255} 와 같습니다.
func (img *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := &Image{256, 256}
	pic.ShowImage(m)
}