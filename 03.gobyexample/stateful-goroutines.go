package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads: // receive (<- chan)
				read.resp <- state[read.key]
			case write := <-writes: // receive (<- chan)
				state[write.key] = write.val
				write.resp <- true // sends (chan <-)
			}
		}
	}()
	fmt.Println("for state doing...")

	for r := 0; r < 100; r++ {
		go func(nIdx int) {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read // sends (chan <-)
				<-read.resp // receive (<- chan)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}(r)
	}
	fmt.Println("for r doing...")

	for w := 0; w < 10; w++ {
		go func(nIdx int) {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write // sends (chan <-)
				<-write.resp // receive (<- chan)
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}(w)
	}
	fmt.Println("for w doing...")

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
