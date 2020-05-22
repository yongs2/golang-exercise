package main

import (
	"log"
	"strings"
)

func main() {
	Example_Count()
	Example_Index()
	Example_ReplaceAll()
	Example_Replace()
	Example_LastIndex()
	Example_IndexCharacterInString()
}

func Example_Count() {
	input := "abcdabcd"
	sub := "ab"
	res := strings.Count(input, sub)
	log.Printf("Count(%v, %v)=%v\n", input, sub, res)

	input = "abcdabcd"
	sub = "xy"
	res = strings.Count(input, sub)
	log.Printf("Count(%v, %v)=%v\n", input, sub, res)

	input = "abcdabcd"
	sub = ""
	res = strings.Count(input, sub)
	log.Printf("Count(%v, %v)=%v\n", input, sub, res)
}

func Example_Index() {
	input := "abcdef"
	sub := "bc"
	res := strings.Index(input, sub)
	log.Printf("Index(%v, %v)=%v\n", input, sub, res)

	input = "abcdefabcdef"
	sub = "cd"
	res = strings.Index(input, sub)
	log.Printf("Index(%v, %v)=%v\n", input, sub, res)

	input = "abcdef"
	sub = "ba"
	res = strings.Index(input, sub)
	log.Printf("Index(%v, %v)=%v\n", input, sub, res)
}

func Example_ReplaceAll() {
	input := "abcdabxyabr"
	old := "ab"
	new := "12"
	res := strings.ReplaceAll(input, old, new)
	log.Printf("ReplaceAll(%v, %v, %v)=%v\n", input, old, new, res)

	input = "abcdabxyabr"
	old = ""
	new = "12"
	res = strings.ReplaceAll(input, old, new)
	log.Printf("ReplaceAll(%v, %v, %v)=%v\n", input, old, new, res)
}

func Example_Replace() {
	input := "abcdabxyabr"
	old := "ab"
	new := "12"
	count := 1
	res := strings.Replace(input, old, new, count)
	log.Printf("Replace(%v, %v, %v, %v)=%v\n", input, old, new, count, res)

	input = "abcdabxyabr"
	old = "ab"
	new = "12"
	count = -1
	res = strings.Replace(input, old, new, count)
	log.Printf("Replace(%v, %v, %v, %v)=%v\n", input, old, new, count, res)

	input = "abcdabxyabr"
	old = "a"
	new = "1"
	count = 1
	res = strings.Replace(input, old, new, count)
	log.Printf("Replace(%v, %v, %v, %v)=%v\n", input, old, new, count, res)

	input = "abcdabxyabr"
	old = "a"
	new = "1"
	count = -1
	res = strings.Replace(input, old, new, count)
	log.Printf("Replace(%v, %v, %v, %v)=%v\n", input, old, new, count, res)
}

func Example_LastIndex() {
	input := "abcdef"
	sub := "bc"
	res := strings.LastIndex(input, sub)
	log.Printf("LastIndex(%v, %v)=%v\n", input, sub, res)

	input = "abcdefabcdef"
	sub = "cd"
	res = strings.LastIndex(input, sub)
	log.Printf("LastIndex(%v, %v)=%v\n", input, sub, res)

	input = "abcdef"
	sub = "ba"
	res = strings.LastIndex(input, sub)
	log.Printf("LastIndex(%v, %v)=%v\n", input, sub, res)
}

func Example_IndexCharacterInString() {
	sample := "ab£c"

	// byte
	for i := 0; i < len(sample); i++ {
        log.Printf("[%d]=[%c]\n", i, sample[i])
	}
	log.Printf("byte.Length is %d\n", len(sample))
	
	sampleRune := []rune(sample)
	for i := 0; i < len(sampleRune); i++ {
		log.Printf("[%d]=[%c]\n", i, sampleRune[i])
	}
	log.Printf("rune.Length is %d\n", len(sampleRune))
}
