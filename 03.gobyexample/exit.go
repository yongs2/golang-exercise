package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	fmt.Println("START...")
	os.Exit(3)
	fmt.Println("END....")
}
