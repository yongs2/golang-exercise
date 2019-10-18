package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one" // sends (chan <-)
	queue <- "two" // sends (chan <-)
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
