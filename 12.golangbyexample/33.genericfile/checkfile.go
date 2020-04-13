package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	info, err := os.Stat("temp")
	if os.IsNotExist(err) {
		log.Fatal("File does not exist")
	}
	if info.IsDir() {
		log.Println("temp is a directory")
		log.Println(info)
	} else {
		log.Println("temp is a file")
		log.Println(info)
	}
}
