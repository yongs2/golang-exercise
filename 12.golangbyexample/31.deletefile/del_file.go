package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deletes sample.txt file from the current working directory")
}
