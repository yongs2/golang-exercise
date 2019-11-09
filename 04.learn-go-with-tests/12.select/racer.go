package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {
	select {
	case <-ping(a) :
		fmt.Println("Racer:", a, time.Now())
		return a, nil
	case <-ping(b) :
		fmt.Println("Racer:", b, time.Now())
		return b, nil
	case <-time.After(10 * time.Millisecond) :
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{}  {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	fmt.Println("Ping:", url, time.Now())
	return ch
}
