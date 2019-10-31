package main

import "fmt"

func Hello(name string) string {
	if name != "" {
		return "Hello, " + name
	}
	return "Hello, world"
}

func main() {
	fmt.Println(Hello(""))
}
