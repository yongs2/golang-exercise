package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	err := ioutil.WriteFile("temp.txt", []byte("first line\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("second line"); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
