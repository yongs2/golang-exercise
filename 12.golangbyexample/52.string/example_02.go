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
	Example_Permutation()
	Example_IndexCharacterInString()
	Example_Swap()
	Example_SwapString()
	Example_Reverse()
	Example_DeleteChar()
	Example_Repeat()
	Example_EqualFold()
	Example_AscII()
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

	input = "abcdabxyabr"
	old = "a"
	new = ""
	res = strings.ReplaceAll(input, old, new)
	log.Printf("ReplaceAll(%v, %v, %v)=%v\n", input, old, new, res)

	input = "abcabcabcdef"
	old = "ab"
	new = ""
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

func Example_Permutation() {
	sample := "ab£"
	sampleRune := []rune(sample)
	generatePermutation(sampleRune, 0, len(sampleRune)-1)
}

func generatePermutation(sampleRune []rune, left, right int) {
	if left == right {
		log.Printf("generatePermutation=[%v]\n", string(sampleRune))
	} else {
		for i := left; i <= right; i++ {
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
			generatePermutation(sampleRune, left+1, right)
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
		}
	}
}

func Example_Swap() {
	sample := "ab£d"
	r := []rune(sample)

	log.Printf("Before %s\n", string(r))
	r[2], r[3] = r[3], r[2]
	log.Printf("After %s\n", string(r))
}

func Example_SwapString() {
	a := "123"
	b := "xyz"

	log.Printf("Before a:%s,b:%s\n", a, b)
	a, b = b, a
	log.Printf("After a:%s,b:%s\n", a, b)
}

func Example_Reverse() {
	sample := "ab£d"
	r := []rune(sample)
	log.Printf("Before %s\n", string(r))
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}
	log.Printf("After %s\n", string(res))
}

func Example_DeleteChar() {
	sample := "ab£c"
	r := []rune(sample)
	log.Printf("Before %s\n", string(r))
	res := delChar(r, 2)
	log.Printf("After %s\n", string(res))
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func Example_Repeat() {
	input := "abc"
	count := 4
	res := strings.Repeat(input, count)
	log.Printf("Repeat(%v, %v)=[%v]\n", input, count, string(res))
}

func Example_EqualFold() {
	left := "abc"
	right := "ABC"
	log.Printf("EqualFold(%v,%v)=[%v]\n", left, right, strings.EqualFold(left, right))

	left = "abc"
	right = "aBC"
	log.Printf("EqualFold(%v,%v)=[%v]\n", left, right, strings.EqualFold(left, right))

	left = "abc"
	right = "AbC"
	log.Printf("EqualFold(%v,%v)=[%v]\n", left, right, strings.EqualFold(left, right))

	left = "abc"
	right = "AbCd"
	log.Printf("EqualFold(%v,%v)=[%v]\n", left, right, strings.EqualFold(left, right))
}

func Example_AscII() {
	lowercase := "abcdefghijklmnopqrstunwxyz"
	for _, c := range lowercase {
		log.Printf("[%v]=[%c]", c, c)
	}

	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range uppercase {
		log.Printf("[%v]=[%c]", c, c)
	}

	numbers := "0123456789"
	for _, n := range numbers {
		log.Printf("[%v]=[%c]", n, n)
	}
}
