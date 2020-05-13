package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go execute(i)
	}
	log.Printf("NumGoroutine=[%v]\n", runtime.NumGoroutine())
}

func execute(idx int) {
	log.Printf("R[%d] Start\n", idx)
	time.Sleep(10 * time.Second)
	log.Printf("R[%d] Done\n", idx)
}
