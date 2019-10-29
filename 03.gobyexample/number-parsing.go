package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(1, f)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(2, i)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(3, d)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(4, u)
	k, _ := strconv.Atoi("135")
	fmt.Println(5, k)
	_, e := strconv.Atoi("wat")
	fmt.Println(6, e)
}
