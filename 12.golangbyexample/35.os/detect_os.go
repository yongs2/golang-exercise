package main

import (
	"log"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "windows":
		log.Println("windows")
	case "darwin":
		log.Println("MAC operating system")
	case "linux":
		log.Println("Linux")
	default:
		log.Printf("%s.\n", os)
	}
}
