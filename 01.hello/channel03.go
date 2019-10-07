package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	fmt.Println("fibonacci.for, n=", n, cap(c))
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	fmt.Println("fibonacci.close(c), n=", n)
	close(c)
}

func main() {
	c := make(chan int, 10)
	fmt.Println("c=", cap(c))
	
	go fibonacci(cap(c), c)
	fmt.Println("---before for...")
	for i := range c {	// 반복문은 채널이 닫힐 때까지 계속해서 값을 받습니다.
		fmt.Println(i)
	}
	fmt.Println("---after for...")
}