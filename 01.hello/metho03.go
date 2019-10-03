package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 포인터 리시버를 사용
// v 를 Vertex 타입으로 받으면 Scale 메소드가 더 이상 동작하지 않습니다.
// Scale 은 v 를 바꾸는데, v 가 (포인터가 아닌) 값 타입이기 때문에 Vertex 타입인 복사본에 작업을 하기 때문에 원래의 값은 바뀌지 않습니다.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 포인터 리시버를 사용
// Abs 의 경우는 다릅니다. 여기서는 v 를 읽기만 하기 때문에, (포인터가 가르키는) 원래의 값이건 복사본이건 상관이 없게 됩니다.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}