package main

import (
	"fmt"
)

/*
chan : bidirectional channel (Both read and write)
chan <- : only writing to channel
<- chan : only reading from the channel (input channel)
*chan : channel pointer. Both read and write
*/

func main() {
	ch := make(chan int, 3)
	ch <- 2 // write to channel
	process(ch)

}

// Only Receive Channel
func process(ch chan int) {
	s := <-ch // read from the channel
	fmt.Println(s)
}
