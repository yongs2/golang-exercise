package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	fmt.Println("fibonacci.for")
	for {
		// select 는 case 구문으로 받는 통신 동작들 중 하나가 수행될 수 있을 때까지 수행을 블록합니다.
		// 다수의 채널이 동시에 준비되면 그 중 하나를 무작위로 선택합니다.
		select {
		case c <- x :	// send x to c
			x, y = y, x+y
		case <- quit :
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("func.start")
		for i :=0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0	// send 0 to quit
		fmt.Println("func.quit")
	}()
	fibonacci(c, quit)
}