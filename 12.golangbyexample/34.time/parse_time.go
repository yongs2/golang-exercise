package main

import (
	"fmt"
	"time"
)

func main() {
	timeT, _ := time.Parse("2006-01-02", "2020-01-29")
	fmt.Println(timeT)

	timeT, _ = time.Parse("06-01-02", "20-01-29")
	fmt.Println(timeT)

	timeT, _ = time.Parse("2006-Jan-02", "2020-Jan-29")
	fmt.Println(timeT)

	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Jan-29 Wednesday 12:19:25")
	fmt.Println(timeT)

	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05 PM MST -07:00", "2020-Jan-29 Wednesday 12:19:25 AM KST +05:30")
	fmt.Println(timeT)
}
