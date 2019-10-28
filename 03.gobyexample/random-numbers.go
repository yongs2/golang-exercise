package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(1, rand.Intn(100), ",")
	fmt.Print(2, rand.Intn(100))
	fmt.Println()

	fmt.Println(3, rand.Float64())

	fmt.Println((rand.Float64()*5)+5, ",")
	fmt.Println((rand.Float64() * 5) + 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(7, r1.Intn(100), ",")
	fmt.Print(8, r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Print(9, r2.Intn(100), ",")
	fmt.Print(10, r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(11, r3.Intn(100), ",")
	fmt.Print(12, r3.Intn(100))
}
