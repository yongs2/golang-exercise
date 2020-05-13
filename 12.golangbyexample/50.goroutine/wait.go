package main

import (
	"log"
	"sync"
	"runtime"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var numOfGoroutine int = 10

	wg.Add(numOfGoroutine)

	for i := 0; i<numOfGoroutine; i++ {
		go sleep(&wg, i, time.Duration(i) * time.Second)
	}
	log.Printf("goroutines[%v]", runtime.NumGoroutine())
	
	wg.Wait()
	log.Printf("[%v] All goroutines finished", numOfGoroutine)
}

func sleep(wg *sync.WaitGroup, idx int, t time.Duration) {
	defer wg.Done()
	time.Sleep(t)
	log.Printf("R[%v] Finished Execution", idx)
}
