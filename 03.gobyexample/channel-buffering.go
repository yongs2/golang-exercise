package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered" // sends (chan <-)
	messages <- "channel"

	fmt.Println("-------")

	fmt.Println(<-messages) // receive (<- chan)
	fmt.Println(<-messages)
}
