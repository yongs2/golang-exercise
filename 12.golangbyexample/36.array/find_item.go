package main

import (
	"log"
)

func main() {
	source := []string{"san", "man", "tan"}

	result := find(source, "san")
	log.Printf("Item %s found: %t\n", "san", result)

	result = find(source, "can")
	log.Printf("Item %s found: %t\n", "can", result)
}

func find(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}

	}
	return false
}
