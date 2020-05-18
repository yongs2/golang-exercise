package main

import (
	"log"
	"reflect"
	"unsafe"
)

func main() {
	var b byte = 'a'
	log.Printf("Priting %v, Size: %d, Type: %s, Character: %c\n", "byte", unsafe.Sizeof(b), reflect.TypeOf(b), b)

	r := '£'
	log.Printf("Priting %v, Size: %d, Type: %s, Unicode CodePoint: %U\n", "rune", unsafe.Sizeof(r), reflect.TypeOf(r), r)

	s := "µ"
	log.Printf("Priting %v, Size: %d, Type: %s, Character: %s\n", "string", unsafe.Sizeof(s), reflect.TypeOf(s), s)
}
