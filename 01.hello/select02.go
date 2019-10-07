package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <- tick :
			fmt.Println("tick.")
		case <- boom :
			fmt.Println("BOOM!")
			return
		default :	// select 의 default 케이스는 현재 수행 준비가 완료된 케이스가 없을 때 수행됩니다.
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}