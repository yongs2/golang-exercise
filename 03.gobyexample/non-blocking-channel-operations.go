package main

import "fmt"

func main() {
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message: // receive (<- chan)
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case message <- msg: // sends (chan <-)
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message: // receive (<- chan)
		fmt.Println("received message:", msg)
	case sig := <-signals: // receive (<- chan)
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}
