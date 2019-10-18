package main

import (
	"fmt"
	"time"
)

// time go run select.go
// real 0m2.293s
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one" // sends (chan <-)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two" // sends (chan <-)
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1: // receive (<- chan)
			fmt.Println("c1.received:", msg1)
		case msg2 := <-c2: // receive (<- chan)
			fmt.Println("c2.received:", msg2)
		}
	}
}
