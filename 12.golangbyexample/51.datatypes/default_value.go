package main

import (
	"log"
)

func main() {
	var a int
	log.Printf("Default zero value of %T: [%v]\n", a, a)

	var b uint
	log.Printf("Default zero value of %T: [%v]\n", b, b)

	var c float32
	log.Printf("Default zero value of %T: [%v]\n", c, c)

	var d complex64
	log.Printf("Default zero value of %T: [%v]\n", d, d)

	var e byte
	log.Printf("Default zero value of %T: [%v], byte\n", e, e)

	var f rune
	log.Printf("Default zero value of %T: [%v], rune\n", f, f)

	var g string
	log.Printf("Default zero value of %T: [%v]\n", g, g)

	var h bool
	log.Printf("Default zero value of %T: [%v]\n", h, h)

	var i [2]bool
	log.Printf("Default zero value of %T: [%v], array\n", i, i)

	type sample struct {
		a int
		b bool
	}
	j := sample{}
	log.Printf("Default zero value of %T: [%v], struct\n", j, j)

	var k map[bool]bool
	log.Printf("Default zero value of %T: [%v], map=%v\n", k, k, k == nil)

	var l chan int
	log.Printf("Default zero value of %T: [%v], channel\n", l, l)

	var m chan interface{}
	log.Printf("Default zero value of %T: [%v], interface\n", m, m)

	var n []int
	log.Printf("Default zero value of %T: [%v], slice=%v\n", n, n, n == nil)

	var o func()
	log.Printf("Default zero value of %T: [%v], function\n", o, o)

	var p *int
	log.Printf("Default zero value of %T: [%v], pointer\n", p, p)
}
