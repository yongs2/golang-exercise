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
	Example_Split()
	Example_Fields()
	Example_Join()
	Example_HasPrefix()
	Example_HasSuffix()
	Example_ToLower()
	Example_ToUpper()
	Example_Title()
	Example_TrimPrefix()
	Example_TrimSuffix()
	Example_TrimSpace()
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

func Example_Split() {
	input := "ab$cd$ef"
	sep := "$"
	res := strings.Split(input, sep)
	log.Printf("Split(%v, %v)=%v\n", input, sep, res)

	sep = "-"
	res = strings.Split(input, sep)
	log.Printf("Split(%v, %v)=%v\n", input, sep, res)

	sep = ""
	res = strings.Split(input, sep)
	log.Printf("Split(%v, %v)=%v\n", input, sep, res)

	input = ""
	sep = ""
	res = strings.Split(input, sep)
	log.Printf("Split(%v, %v)=%v\n", input, sep, res)
}

func Example_Fields() {
	input := "ab cd ef"
	res := strings.Fields(input)
	log.Printf("Fields(%v)=%v\n", input, res)

	input = "abcdef"
	res = strings.Fields(input)
	log.Printf("Fields(%v)=%v\n", input, res)

	input = "ab  cd   ef "
	res = strings.Fields(input)
	log.Printf("Fields(%v)=%v\n", input, res)

	input = "   "
	res = strings.Fields(input)
	log.Printf("Fields(%v)=%v\n", input, res)

	input = ""
	res = strings.Fields(input)
	log.Printf("Fields(%v)=%v\n", input, res)
}

func Example_Join() {
	input := []string{"ab", "cd", "ef"}
	sep := "-"
	res := strings.Join(input, sep)
	log.Printf("Join(%v, %v)=%v\n", input, sep, res)

	input = []string{}
	sep = ""
	res = strings.Join(input, sep)
	log.Printf("Join(%v, %v)=%v\n", input, sep, res)

	input = []string{"ab", "cd", "ef"}
	sep = ""
	res = strings.Join(input, sep)
	log.Printf("Join(%v, %v)=%v\n", input, sep, res)
}

func Example_HasPrefix() {
	input := "abcdef"
	prefix := "ab"
	res := strings.HasPrefix(input, prefix)
	log.Printf("HasPrefix(%v, %v)=%v\n", input, prefix, res)

	input = "abcdef"
	prefix = "ac"
	res = strings.HasPrefix(input, prefix)
	log.Printf("HasPrefix(%v, %v)=%v\n", input, prefix, res)
}

func Example_HasSuffix() {
	input := "abcdef"
	suffix := "ef"
	res := strings.HasSuffix(input, suffix)
	log.Printf("HasSuffix(%v, %v)=%v\n", input, suffix, res)

	input = "abcdef"
	suffix = "fg"
	res = strings.HasSuffix(input, suffix)
	log.Printf("HasSuffix(%v, %v)=%v\n", input, suffix, res)
}

func Example_ToLower() {
	input := "ABC"
	res := strings.ToLower(input)
	log.Printf("ToLower(%v)=%v\n", input, res)

	input = "ABC12$a"
	res = strings.ToLower(input)
	log.Printf("ToLower(%v)=%v\n", input, res)
}

func Example_ToUpper() {
	input := "abc"
	res := strings.ToUpper(input)
	log.Printf("ToUpper(%v)=%v\n", input, res)

	input = "abc12$"
	res = strings.ToUpper(input)
	log.Printf("ToUpper(%v)=%v\n", input, res)
}

func Example_Title() {
	input := "this is a test sentence"
	res := strings.Title(input)
	log.Printf("Title(%v)=%v\n", input, res)
}

func Example_TrimPrefix() {
	input := "testremoved"
	prefix := "test"
	res := strings.TrimPrefix(input, prefix)
	log.Printf("TrimPrefix(%v, %v)=%v\n", input, prefix, res)

	input = "tesremoved"
	prefix = "test"
	res = strings.TrimPrefix(input, prefix)
	log.Printf("TrimPrefix(%v, %v)=%v\n", input, prefix, res)
}

func Example_TrimSuffix() {
	input := "removedtest"
	prefix := "test"
	res := strings.TrimSuffix(input, prefix)
	log.Printf("TrimSuffix(%v, %v)=%v\n", input, prefix, res)

	input = "removedtes"
	prefix = "test"
	res = strings.TrimSuffix(input, prefix)
	log.Printf("TrimSuffix(%v, %v)=%v\n", input, prefix, res)
}

func Example_TrimSpace() {
	input := " test  "
	res := strings.TrimSpace(input)
	log.Printf("TrimSpace(%v)=%v\n", input, res)

	input = " This is test  "
	res = strings.TrimSpace(input)
	log.Printf("TrimSpace(%v)=%v\n", input, res)
}
