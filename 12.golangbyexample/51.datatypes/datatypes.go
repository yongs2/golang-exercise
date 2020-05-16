package main

import (
	"log"
	"math/bits"
	"reflect"
	"unsafe"
)

type sample struct {
	a int
	b string
}

func main() {
	sizeOfIntInBits := bits.UintSize
	log.Printf("%d bits\n", sizeOfIntInBits)

	var a int
	log.Printf("%s is %d bytes, type is %s\n", "a", unsafe.Sizeof(a), reflect.TypeOf(a))

	b := 2
	log.Printf("%s is %d bytes, type is %s\n", "b", unsafe.Sizeof(b), reflect.TypeOf(b))

	var c int8 = 2
	log.Printf("%s is %d bytes, type is %s\n", "c", unsafe.Sizeof(c), reflect.TypeOf(c))

	var d int16 = 2
	log.Printf("%s is %d bytes, type is %s\n", "d", unsafe.Sizeof(d), reflect.TypeOf(d))

	var e int32 = 2
	log.Printf("%s is %d bytes, type is %s\n", "e", unsafe.Sizeof(e), reflect.TypeOf(e))

	var f int64 = 2
	log.Printf("%s is %d bytes, type is %s\n", "f", unsafe.Sizeof(f), reflect.TypeOf(f))

	var g uint = 2
	log.Printf("%s is %d bytes, type is %s\n", "g", unsafe.Sizeof(g), reflect.TypeOf(g))

	h := &sample{a: 1, b: "test"}
	pointer := unsafe.Pointer(uintptr(unsafe.Pointer(h)) + unsafe.Offsetof(h.b))
	log.Printf(*(*string)(pointer))

	var i uint8 = 2
	log.Printf("%s is %d bytes, type is %s\n", "i", unsafe.Sizeof(i), reflect.TypeOf(i))

	var j uint16 = 2
	log.Printf("%s is %d bytes, type is %s\n", "j", unsafe.Sizeof(j), reflect.TypeOf(j))

	var k uint32 = 2
	log.Printf("%s is %d bytes, type is %s\n", "k", unsafe.Sizeof(k), reflect.TypeOf(k))

	var l uint64 = 2
	log.Printf("%s is %d bytes, type is %s\n", "l", unsafe.Sizeof(l), reflect.TypeOf(l))

	var m float32 = 2
	log.Printf("%s is %d bytes, type is %s\n", "m", unsafe.Sizeof(m), reflect.TypeOf(m))

	var n float64 = 2
	log.Printf("%s is %d bytes, type is %s\n", "n", unsafe.Sizeof(n), reflect.TypeOf(n))

	o := 2.3
	log.Printf("%s is %d bytes, type is %s\n", "o", unsafe.Sizeof(o), reflect.TypeOf(o))

	var p float32 = 3
	var q float32 = 5
	r := complex(p, q)
	log.Printf("%s is %d bytes, type is %s\n", "r", unsafe.Sizeof(r), reflect.TypeOf(r))

	var s complex64
	s = 4 + 5i
	log.Printf("%s is %d bytes, type is %s\n", "s", unsafe.Sizeof(s), reflect.TypeOf(s))

	log.Println(r+s, r-s, r*s, r/s)

	var t float64 = 3
	var u float64 = 5
	v := complex(t, u)
	log.Printf("%s is %d bytes, type is %s\n", "v", unsafe.Sizeof(v), reflect.TypeOf(v))

	w := 4 + 5i
	log.Printf("%s is %d bytes, type is %s\n", "w", unsafe.Sizeof(w), reflect.TypeOf(w))
	log.Println(v+w, v-w, v*w, v/w)

	var x byte = 'a'
	log.Printf("%s is %d bytes, type is %s, Character:%c, byte:%v\n", "x", unsafe.Sizeof(x), reflect.TypeOf(x), x, []byte("abc"))

	y := "0b£"
	log.Printf("UTF-8, %s is %U, %v\n", y, []rune(y), []rune(y))

	z := '£'
	log.Printf("Size:%d, type is %s, Unicode codepoint: %U, Character:%c\n", unsafe.Sizeof(z), reflect.TypeOf(z), z, z)

	z1 := 'a'
	log.Printf("Size:%d, type is %s, Unicode codepoint: %U, Character:%c\n", unsafe.Sizeof(z1), reflect.TypeOf(z1), z1, z1)

	a1 := "this\nthat"
	log.Printf("a1 is: %s\n", a1)

	b1 := `this\nthat`
	log.Printf("b1 is: %s\n", b1)

	c1 := "0b£"
	log.Printf("c1 is: %s, byte:%v, len:%d\n", c1, []byte(c1), len(c1))

	for _, c11 := range c1 {
		log.Printf("c11 is: %s\n", string(c11))
	}

	var d1 bool
	log.Printf("d1's value is: %t\n", d1)

	andOperation := 1 < 2 && 1 > 3
	log.Printf("Ouput of AND operation on one true and other false: %t\n", andOperation)

	orOperation := 1 < 2 || 1 > 3
	log.Printf("Ouput of OR operation on one true and other false: %t\n", orOperation)

	negationOperation := !(1 > 2)
	log.Printf("Ouput of NEGATION operation on one true and other false: %t\n", negationOperation)

	sampleArray := [3]string{"a", "b", "c"}
	log.Printf("Array: %v\n", sampleArray)

	// struct
	type employee struct {
		name   string
		age    int
		salary float64
	}
	employee1 := employee{"John", 21, 1000}
	log.Printf("Struct.%s: %v\n", "employee1", employee1)

	employee2 := employee{
		name:   "Sam",
		age:    22,
		salary: 1100,
	}
	log.Printf("Struct.%s: %v\n", "employee2", employee2)

	employee3 := employee{name: "Tina", age: 24}
	log.Printf("Struct.%s: %v\n", "employee3", employee3)

	// slice
	slice1 := make([]string, 2, 3)
	log.Printf("Slice.%s: %v\n", "slice1", slice1)

	slice2 := []string{"a", "b", "c"}
	log.Printf("Slice.%s: %v\n", "slice2", slice2)

	slice2 = append(slice2, "d")
	log.Printf("Slice.%s: %v\n", "slice2", slice2)

	for _, val := range slice2 {
		log.Printf("Slice.%s.item: %v\n", "slice2", val)
	}
}
