package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			fmt.Println("Any:", v, "true")
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			fmt.Println("All:", v, "false")
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum", "banana"}

	fmt.Println("Index:", Index(strs, "pear"))
	fmt.Println("Include:", Include(strs, "grape"))
	fmt.Println("Any:", Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println("All:", All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println("Filter:", Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println("Map:", Map(strs, strings.ToUpper))
}
