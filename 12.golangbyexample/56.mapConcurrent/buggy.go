package main

import (
	"log"
	"time"
)

var (
	allData = make(map[string]string)
)

func get(key string) string {
	log.Printf("get(%v)=%v\n", key, allData[key])
	return allData[key]
}

func set(key string, value string) {
	allData[key] = value
	log.Printf("set(%v)=%v\n", key, allData[key])
}

func main() {
	go set("a", "Some Data 1")
	go set("b", "Some Data 2")
	go get("a")
	go get("b")
	go get("a")

	time.Sleep(1 * time.Second)
}
