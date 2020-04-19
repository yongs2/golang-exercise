package main

import (
	"fmt"
	"time"
)

func main() {
	tNow := time.Now()

	tUnix := tNow.Unix()
	fmt.Printf("timeUnix %d\n", tUnix)

	timeT := time.Unix(tUnix, 0)
	fmt.Printf("time.Time: %s\n", timeT)
}
