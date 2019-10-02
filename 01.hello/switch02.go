package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("when's saturday?")

	today := time.Now().Weekday()

	// 스위치의 각 조건은 위에서 아래로 평가
	switch time.Saturday {
	case today + 0 :
		fmt.Println("Today.")
	case today + 1 :
		fmt.Println("Tomorrow.")
	case today + 2 :
		fmt.Println("In two days.")
	default :
		fmt.Println("Too far away.", int(time.Saturday - today))
	}
}