package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func main() {
	// 고루틴은 Go 런타임에 의해 관리되는 경량 쓰레드입니다.
	go say("world")
	say("hello")	// 여기서 함수 호출하지 않으면 위의 go 호출은 동작 안함
}
