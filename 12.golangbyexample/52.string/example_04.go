package main

import (
	"log"
)

func main() {
	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	log.Println(string(colorRed), "test")
	log.Println(string(colorGreen), "test")
	log.Println(string(colorYellow), "test")
	log.Println(string(colorBlue), "test")
	log.Println(string(colorPurple), "test")
	log.Println(string(colorWhite), "test")
	log.Println(string(colorCyan), "test", string(colorReset))
	log.Println("next")
}
