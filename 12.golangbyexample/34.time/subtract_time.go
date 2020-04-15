package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	newT := t.Add(-time.Hour * 1)
	fmt.Printf("%32s : %s\n", "Subtracting 1 hour", newT)

	newT = t.Add(-time.Minute * 15)
	fmt.Printf("%32s : %s\n", "Subtracting 15 minute", newT)

	newT = t.Add(-time.Second * 10)
	fmt.Printf("%32s : %s\n", "Subtracting 10 second", newT)

	newT = t.Add(-time.Millisecond * 100)
	fmt.Printf("%32s : %s\n", "Subtracting 100 millisecond", newT)

	newT = t.Add(-time.Millisecond * 1000)
	fmt.Printf("%32s : %s\n", "Subtracting 1000 millisecond", newT)

	newT = t.Add(-time.Nanosecond * 10000)
	fmt.Printf("%32s : %s\n", "Subtracting 10000 Nanosecond", newT)

	newT = t.AddDate(-1, -2, -4)
	fmt.Printf("%32s : %s\n", "Subtracting 1 year 2 month 4 day", newT)
}
