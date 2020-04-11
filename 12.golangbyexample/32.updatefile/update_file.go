package main

import (
	"log"
	"os"
	"time"
)

func main() {
	fileName := "sample.txt"

	// Create
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CreateFile=[%v]\n", file)

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ModTime=[%v]\n", fileInfo.ModTime())

	time.Sleep(time.Second)

	// Update
	currentTime := time.Now().Local()
	err = os.Chtimes(fileName, currentTime, currentTime)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Changes the atime and mtime of the file.")

	fileInfo, err = os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ModTime=[%v]\n", fileInfo.ModTime())

	// Remove
	err = os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deletes %s file from the current working directory\n", fileName)
}
