package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	oldTime := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	diff := currentTime.Sub(oldTime)

	fmt.Printf("Hours: %f\n", diff.Hours())

	fmt.Printf("Minutes: %f\n", diff.Minutes())

	fmt.Printf("Seconds: %f\n", diff.Seconds())

	fmt.Printf("Nanoseconds: %d\n", diff.Nanoseconds())

	fmt.Println(time.Since(currentTime.Add(-time.Hour * 1)))

	fmt.Println(time.Until(currentTime.Add(time.Hour * 1)))
}
