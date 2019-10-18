package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg // sends (chan <-)
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // receive (<- chan)
	pongs <- msg   // sends (chan <-)
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs) // receive (<- chan)
}
