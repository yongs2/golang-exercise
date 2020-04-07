package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	WordByWordScan()
}

func WordByWordScan() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// bufio.Scanner has max buffer size 64*1024 bytes
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
