package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // receive (<- chan)
			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true // sends (chan <-)
				return
			}
		}
	}()

	for j := 1; j < 8; j++ {
		jobs <- j // sends (chan <-)
		fmt.Println("Send jobs:", j)
	}

	// close
	close(jobs)
	fmt.Println("Send All jobs")

	<-done // receive (<- chan)
}
