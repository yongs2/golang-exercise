package main

import (
	"log"
	"unsafe"
)

type sample struct {
	a int
	b string
}

func main() {
	s := &sample{a: 1, b: "test"}

	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))
	log.Printf("unsafe.Pointer(s)=[%v]\n", unsafe.Pointer(s))
	log.Printf("unsafe.Offsetof(s.a)=[%v]\n", unsafe.Offsetof(s.a))
	log.Printf("unsafe.Offsetof(s.b)=[%v]\n", unsafe.Offsetof(s.b))
	log.Printf("*string=[%v]\n", *(*string)(p))

	startAddress := uintptr(unsafe.Pointer(s))
	log.Printf("Start Address of s: %v\n", startAddress)
}
