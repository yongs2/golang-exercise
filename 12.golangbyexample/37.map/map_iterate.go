package main

import (
	"log"
	"time"
)

func main() {
	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	// Iterating over all keys and values
	start := time.Now()
	for k, v := range sample {
		log.Printf("key: %s, Value: %s\n", k, v)
	}
	log.Printf("Elapsed: %v\n", time.Since(start))

	// Iterating over only keys
	start = time.Now()
	for k := range sample {
		log.Printf("Keys: %s\n", k)
	}
	log.Printf("Elapsed: %v\n", time.Since(start))

	// Iterating over only keys
	result := ""
	start = time.Now()
	for k := range sample {
		result += k + " "
	}
	log.Printf("Keys: %s\n", result)
	log.Printf("Elapsed: %v\n", time.Since(start))

	// Iterating over only values
	start = time.Now()
	for _, v := range sample {
		log.Printf("Value: %s\n", v)
	}
	log.Printf("Elapsed: %v\n", time.Since(start))

	// Get list of all keys
	start = time.Now()
	keys := getAllKeys(sample)
	log.Printf("Keys: %s\n", keys)
	log.Printf("Elapsed: %v\n", time.Since(start))
}

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	return keys
}
