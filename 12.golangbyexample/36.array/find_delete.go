package main

import (
	"log"
)

func main() {
	log.Println("---------------")
	// Keep copying to the original slice
	before1 := []int{1, 2, 3, 1}
	log.Println("before:", before1)
	after1 := findAndDelete_Copy(before1, 1)
	log.Println("after:", after1)
	log.Println("before:", before1)

	log.Println("---------------")
	// Create a new slice and keep inserting into it
	before2 := []int{1, 2, 3, 1}
	log.Println("before:", before2)
	after2 := findAndDelete_New(before2, 1)
	log.Println("after:", after2)
	log.Println("before:", before2)
}

func findAndDelete_Copy(s []int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}

func findAndDelete_New(s []int, itemToDelete int) []int {
	var new = make([]int, 0, len(s))
	index := 0
	for _, i := range s {
		if i != itemToDelete {
			new = append(new, i)
			index++
		}
	}
	return new[:index]
}
