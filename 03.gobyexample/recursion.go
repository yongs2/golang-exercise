package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		fmt.Println("fact(", n, ")")
		return 1
	}
	fmt.Println("fact(", n, ")")
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(7))
}
