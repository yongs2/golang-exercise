package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // receive (<- timer)
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C // receive (<- timer)
		fmt.Println("Timer2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}
