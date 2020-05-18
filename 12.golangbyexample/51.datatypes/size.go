package main

import (
	"log"
	"math/bits"
	"unsafe"
)

func main() {
	sizeInBits := bits.UintSize
	log.Printf("%d bits\n", sizeInBits)

	var a int
	log.Printf("%d bytes\n", unsafe.Sizeof(a))

	var b uint
	log.Printf("%d bytes\n", unsafe.Sizeof(b))
}
