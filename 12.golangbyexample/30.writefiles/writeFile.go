package main

import (
	"io/ioutil"
	"log"
)

func main() {
	linesToWrite := "This is an example to show how to write to file using ioutil."
	err := ioutil.WriteFile("temp.txt", []byte(linesToWrite), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
