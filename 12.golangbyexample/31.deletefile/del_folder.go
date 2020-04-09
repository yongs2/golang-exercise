package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("sample")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deletes sample folder from the current working directory")
}
