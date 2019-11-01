package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) ([]int) {
	var sums []int
	
	for i, numbers := range numbersToSum {
		fmt.Println("Idx=", i, "numbers=", numbers)
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTail(numbersToSum ...[]int) ([]int) {
	var sums []int
	
	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			fmt.Println("Idx=", i, "tail=", tail)
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
