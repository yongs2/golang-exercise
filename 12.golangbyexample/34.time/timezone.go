package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("%20s: %s\n", "UTC Time", now.In(loc))

	loc, _ = time.LoadLocation("Europe/Berlin")
	fmt.Printf("%20s: %s\n", "Berlin Time", now.In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("%20s: %s\n", "New York Time", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Dubai")
	fmt.Printf("%20s: %s\n", "Dubai Time", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Seoul")
	fmt.Printf("%20s: %s\n", "Seoul Time", now.In(loc))
}
