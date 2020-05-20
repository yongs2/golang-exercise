package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	Example_strconv()
	Example_ReplaceAll()
	Example_multiline()
	Exmaple_Compare()
	Exmaple_Contains()
}

func Example_strconv() {
	x := "1234"
	val, err := strconv.Atoi(x)
	if err != nil {
		log.Printf("Supplied value %s is not a number\n", x)
	} else {
		log.Printf("Supplied value %s is a number with value %d\n", x, val)
	}

	y := "123b"
	val, err = strconv.Atoi(y)
	if err != nil {
		log.Printf("Supplied value %s is not a number\n", y)
	} else {
		log.Printf("Supplied value %s is a number with value %d\n", y, val)
	}
}

func Example_ReplaceAll() {
	sample := " This is a sample string   "
	noSpaceString := strings.ReplaceAll(sample, " ", "")
	log.Printf("Org:[%v], ReplaceAll: [%v]\n", sample, noSpaceString)
}

func Example_multiline() {
	multiline := `This is 
a multiline 
string`
	log.Printf("Multiline: [%v]\n", multiline)
}

func Exmaple_Compare() {
	res := strings.Compare("abc", "abc")
	log.Printf("Compare(%v, %v)=%v\n", "abc", "abc", res)

	res = strings.Compare("abc", "xyz")
	log.Printf("Compare(%v, %v)=%v\n", "abc", "xyz", res)

	res = strings.Compare("xyz", "abc")
	log.Printf("Compare(%v, %v)=%v\n", "xyz", "abc", res)
}

func Exmaple_Contains() {
	present := strings.Contains("abc", "ab")
	log.Printf("Contains(%v, %v)=%v\n", "abc", "ab", present)

	present = strings.Contains("abc", "xyz")
	log.Printf("Contains(%v, %v)=%v\n", "abc", "xyz", present)
}
