package main

import (
	"log"
	"sync"
)

var (
	allData = make(map[string]string)
	rwm     sync.RWMutex
)

func get(key string) string {
	rwm.RLock()
	defer rwm.RUnlock()
	return allData[key]
}

func set(key string, value string) {
	rwm.Lock()
	defer rwm.Unlock()
	allData[key] = value
}

func main() {
	set("a", "Some Data")
	result := get("a")
	log.Println(result)
}
