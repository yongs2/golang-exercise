package main

import (
	"log"
	"sync"
	"time"
)

func f(w *sync.WaitGroup, sec int) {
	log.Println("a -----> ", sec)
	time.Sleep(time.Duration(sec) * time.Second)
	log.Println("b <----- ", sec)
	w.Done()
}

func main() {
	log.Println("start")
	var w sync.WaitGroup
	w.Add(2)

	go func(w *sync.WaitGroup, msec int) {
		log.Println("c -----> ", msec)
		<-time.After(time.Duration(msec) * time.Millisecond)
		log.Println("d <----- ", msec)
		w.Done()
	}(&w, 2500)

	go f(&w, 4)

	w.Wait()

	time.Sleep(900000000 * time.Nanosecond)
	log.Println("finish")
}
