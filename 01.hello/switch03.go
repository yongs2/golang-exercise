package main

import ( 
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// 스위치에서 조건을 생략하면 " switch true " 와 같습니다.
	switch {
	case t.Hour() < 12 :
		fmt.Println("Good Morning")
	case t.Hour() < 17 :
		fmt.Println("Good Afternoon")
	default :
		fmt.Println("Good evening")
	}
}