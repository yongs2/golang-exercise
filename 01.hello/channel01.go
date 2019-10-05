package main

import "fmt"

// 채널은 채널 연산자 <- 를 이용해 값을 주고 받을 수 있는, 타입이 존재하는 파이프입니다.
func sum(a []int, c chan int) {
	sum := 0
	fmt.Println(a)
	for _, v := range a {
		sum += v
	}
	fmt.Println(a, sum)
	c <- sum 	// send sum to c
}

func main() {
	a := [] int {7, 2, 8, -9, 4, 0}

	// 맵이나 슬라이스처럼, 채널은 사용되기 전에 생성되어야 합니다.
	// 송/수신은 상대편이 준비될 때까지 블록됩니다.
	// 고루틴이 명시적인 락이나 조건 없이도 동기화 
	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c 	// receive from c

	fmt.Println(x, y, x+y)
}