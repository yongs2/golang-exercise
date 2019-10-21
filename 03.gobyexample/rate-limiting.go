package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i // sends (chan <-)
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter // receive (<- chan) // 200 ms 단위로 blocking됨
		fmt.Println("1.request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now() // sends (chan <-)
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t // sends (chan <-)
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i // sends (chan <-)
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter // receive (<- chan), burstyLimiter가 3개이어서, 3개만 blocking 만 되고, 그외는 blocking안됨
		fmt.Println("2.request", req, time.Now())
	}
}
