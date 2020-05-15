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
	var u float64 = 3
	v := complex(t, u)
	log.Printf("%s is %d bytes, type is %s\n", "v", unsafe.Sizeof(v), reflect.TypeOf(v))

	w := 4 + 5i
	log.Printf("%s is %d bytes, type is %s\n", "w", unsafe.Sizeof(w), reflect.TypeOf(w))
	log.Println(v-w, v-w, v*w, v/w)

	var x byte = 'a'
	log.Printf("%s is %d bytes, type is %s\n", "x", unsafe.Sizeof(w), reflect.TypeOf(w))
}
