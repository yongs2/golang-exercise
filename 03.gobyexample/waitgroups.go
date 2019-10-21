package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
	wg.Done() // Notify the WaitGroup that this worker is done.
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment the WaitGroup counter
		go worker(i, &wg)
	}
	wg.Wait() // Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	fmt.Println("DONE...")
}
