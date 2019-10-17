package main

import "fmt"

func main() {
	messages := make(chan string)

	fmt.Println("before")
	go func() {
		messages <- "ping"
	}()
	fmt.Println("after")

	msg := <-messages
	fmt.Println(msg)
}
