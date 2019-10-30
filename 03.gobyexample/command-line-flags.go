package main

import (
	"flag"
	"fmt"
)

// go build command-line-flags.go
// ./command-line-flags -word=opt -numb=7 -fork -svar=flag
func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "a int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}