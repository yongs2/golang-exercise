package main

import "fmt"

func main() {
	// make 에 두번째 인자로 버퍼 용량을 넣음으로써 해당 용량만큼 버퍼링되는 채널을 생성할 수 있습니다:
	var max int = 100

	c := make(chan int, max)

	// 버퍼링되는 채널로의 송신은 버퍼가 꽉 찰 때까지 블록됩니다. 수신측은 버퍼가 빌 때 블록됩니다.	
	for i := 0; i < max; i++ {
		c <- i
	}

	for i := 0; i < max; i++ {
		fmt.Println(<-c)
	}
}