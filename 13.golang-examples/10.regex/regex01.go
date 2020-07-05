package main

import (
	"log"
	"regexp"
)

var variable string
var strArray []string
var r = regexp.MustCompile("[^\\s]+")

func main() {
	variable = "Lorem  Ipsum  Dolor   Sit  Amet"
	strArray = r.FindAllString(variable, -1)
	for i := 0; i < len(strArray); i++ {
		log.Println(strArray[i])
	}
}
