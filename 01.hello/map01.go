package main

import "fmt"

type Vertex struct {
	// Latitude of the event
	// Longitude of the event
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// 맵은 반드시 사용하기 전에 make 를 명시
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{ 
		40.68433, -74.39967,	// } 를 다음 라인으로 하려면, "," 를 추가해야 함
	}
	fmt.Println(m["Bell Labs"])
}