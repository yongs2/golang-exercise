package sort

import (
	"math/rand"
)

func generateRandomNumber(size int) []int {
	rand_number := make([]int, size, size)
	for i := 0; i < size; i++ {
		rand_number[i] = rand.Intn(size)
	}

	return rand_number
}
