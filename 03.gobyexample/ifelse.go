package main

import "fmt"

func main() {
	var i int = 7
	if i%2 == 0 {
		fmt.Println(i, "is even")
	} else {
		fmt.Println(i, "is odd")
	}

	var j, k int = 8, 4
	if j%k == 0 {
		fmt.Println(j, "is divisible by", k)
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digit")
	}
}
