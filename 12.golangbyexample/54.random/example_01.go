package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	Example_rand()
	RandomPassword()
	example_PickNumRandomInSlice()
	example_PickCharRandomInSlice()
	example_ShuffleString()
	example_ShuffleSlice()
	example_Perm()
	example_RandomNumber()
	example_RandomString()
	example_RandomRune()
}

func Example_rand() {
	var val int

	rand.Seed(time.Now().Unix())
	val = 10
	log.Printf("rand.Intn(%v)=%v", val, rand.Intn(val))

	val = 10
	log.Printf("rand.Int31n(%v)=%v", val, rand.Int31n(int32(val)))

	val = 10
	log.Printf("rand.Int63n(%v)=%v", val, rand.Int63n(int64(val)))

	log.Printf("rand.Int()=%v", rand.Int())
	log.Printf("rand.Int31()=%v", rand.Int31())
	log.Printf("rand.Int63()=%v", rand.Int63())
	log.Printf("rand.Uint32()=%v", rand.Uint32())
	log.Printf("rand.Uint64()=%v", rand.Uint64())
	log.Printf("rand.Float64()=%v", rand.Float64())
	log.Printf("rand.Float32()=%v", rand.Float32())
}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func RandomPassword() {
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	log.Printf("generatePassword=[%v]", password)

	minSpecialChar = 2
	minNum = 2
	minUpperCase = 2
	passwordLength = 20
	password = generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	log.Printf("generatePassword=[%v]", password)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func example_PickNumRandomInSlice() {
	in := []int{2, 5, 6}
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(in))
		pick := in[randomIndex]
		log.Printf("[%v] : in[%v]=%v", i, randomIndex, pick)
	}
}

func example_PickCharRandomInSlice() {
	in := "abcdedf£"
	inRune := []rune(in)
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(len(inRune))
		pick := inRune[randomIndex]
		log.Printf("[%v] : inRune[%v]=%v", i, randomIndex, string(pick))
	}
}

func example_ShuffleString() {
	rand.Seed(time.Now().Unix())

	in := "abcdedf£"
	inRune := []rune(in)

	for i := 0; i < 10; i++ {
		rand.Shuffle(len(inRune), func(i, j int) {
			inRune[i], inRune[j] = inRune[j], inRune[i]
		})
		log.Printf("[%v] : shuffle=%v", i, string(inRune))
	}
}

func example_ShuffleSlice() {
	rand.Seed(time.Now().Unix())

	in := []int{2, 3, 5, 8}

	for i := 0; i < 10; i++ {
		rand.Shuffle(len(in), func(i, j int) {
			in[i], in[j] = in[j], in[i]
		})
		log.Printf("[%v] : shuffle=%v", i, in)
	}
}

func example_Perm() {
	rand.Seed(time.Now().Unix())
	var val int

	for i := 0; i < 10; i++ {
		val = 10
		log.Printf("Perm(%v)=%v", val, rand.Perm(val))
	}
}

func example_RandomNumber() {
	rand.Seed(time.Now().Unix())
	var rangeLower, rangeUpper, randomNum int

	rangeLower = 5
	rangeUpper = 20
	for i := 0; i < 10; i++ {
		randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
		log.Printf("[%v]: [%v] <= [%v] <= [%v]", i, rangeLower, randomNum, rangeUpper)
	}

	rangeLower = 100
	rangeUpper = 200
	for i := 0; i < 10; i++ {
		randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
		log.Printf("[%v]: [%v] <= [%v] <= [%v]", i, rangeLower, randomNum, rangeUpper)
	}
}

func example_RandomString() {
	rand.Seed(time.Now().Unix())

	var charSet string
	var output strings.Builder
	var length, random int
	var randomChar byte

	charSet = "abcdedfghijklmnopqrst"
	length = 10
	for i := 0; i < length; i++ {
		random = rand.Intn(len(charSet))
		randomChar = charSet[random]
		output.WriteString(string(randomChar))
	}
	log.Printf("WriteString=[%v]", output.String())
	output.Reset()

	charSet = "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length = 20
	for i := 0; i < length; i++ {
		random = rand.Intn(len(charSet))
		randomChar = charSet[random]
		output.WriteString(string(randomChar))
	}
	log.Printf("WriteString=[%v]", output.String())
}

func example_RandomRune() {
	rand.Seed(time.Now().Unix())

	var charSet []rune
	var output strings.Builder
	var length, random int
	var randomChar rune

	charSet = []rune("abcdedfghijklmnopqrst£")
	length = 10
	for i := 0; i < length; i++ {
		random = rand.Intn(len(charSet))
		randomChar = charSet[random]
		output.WriteRune(randomChar)
	}
	log.Printf("WriteRune=[%v]", output.String())
	output.Reset()

	charSet = []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP£")
	length = 20
	for i := 0; i < length; i++ {
		random = rand.Intn(len(charSet))
		randomChar = charSet[random]
		output.WriteRune(randomChar)
	}
	log.Printf("WriteRune=[%v]", output.String())
}
