package main

import (
	"log"
	"math"
)

func main() {
	example_Log()
	example_Logb()
	example_Log2()
	example_Log10()
	example_ModulusOperator()
	example_Mod()
	example_Remainder()
	example_Modf()
	example_Pow()
	example_Signbit()
	example_Min()
	example_Max()
}

func example_Log() {
	var f, res float64

	f = 4
	res = math.Log(f)
	log.Printf("Log(%v)=[%v]", f, res)

	f = 10.2
	res = math.Log(f)
	log.Printf("Log(%v)=[%v]", f, res)

	f = -10
	res = math.Log(f)
	log.Printf("Log(%v)=[%v]", f, res)
}

func example_Logb() {
	var f, res float64

	f = 4
	res = math.Logb(f)
	log.Printf("Logb(%v)=[%v]", f, res)

	f = 10.2
	res = math.Logb(f)
	log.Printf("Logb(%v)=[%v]", f, res)

	f = -10
	res = math.Logb(f)
	log.Printf("Logb(%v)=[%v]", f, res)
}

func example_Log2() {
	var f, res float64

	f = 4
	res = math.Log2(f)
	log.Printf("Log2(%v)=[%v]", f, res)

	f = 10.2
	res = math.Log2(f)
	log.Printf("Log2(%v)=[%v]", f, res)

	f = -10
	res = math.Log2(f)
	log.Printf("Log2(%v)=[%v]", f, res)
}

func example_Log10() {
	var f, res float64

	f = 100
	res = math.Log10(f)
	log.Printf("Log10(%v)=[%v]", f, res)

	f = 10
	res = math.Log10(f)
	log.Printf("Log10(%v)=[%v]", f, res)

	f = -10
	res = math.Log10(f)
	log.Printf("Log10(%v)=[%v]", f, res)
}

func example_ModulusOperator() {
	var a, b, res int

	a = 4
	b = 2
	res = a % b
	log.Printf("%v %% %v =[%v]", a, b, res)

	a = 5
	b = 2
	res = a % b
	log.Printf("%v %% %v =[%v]", a, b, res)

	a = 8
	b = 3
	res = a % b
	log.Printf("%v %% %v =[%v]", a, b, res)
}

func example_Mod() {
	var a, b, res float64

	a = 4
	b = 2
	res = math.Mod(a, b)
	log.Printf("Mod(%v, %v)=[%v]", a, b, res)

	a = 4.2
	b = 2
	res = math.Mod(a, b)
	log.Printf("Mod(%v, %v)=[%v]", a, b, res)

	a = 5
	b = 2
	res = math.Mod(a, b)
	log.Printf("Mod(%v, %v)=[%v]", a, b, res)

	a = -5
	b = 2
	res = math.Mod(a, b)
	log.Printf("Mod(%v, %v)=[%v]", a, b, res)
}

func example_Remainder() {
	var a, b, res float64

	a = 4
	b = 2
	res = math.Remainder(a, b)
	log.Printf("Remainder(%v, %v)=[%v]", a, b, res)

	a = 5
	b = 2
	res = math.Remainder(a, b)
	log.Printf("Remainder(%v, %v)=[%v]", a, b, res)

	a = 5.5
	b = 2
	res = math.Remainder(a, b)
	log.Printf("Remainder(%v, %v)=[%v]", a, b, res)

	a = 5.5
	b = 1.5
	res = math.Remainder(a, b)
	log.Printf("Remainder(%v, %v)=[%v]", a, b, res)
}

func example_Modf() {
	var f, integer, fraction float64

	f = 2.5
	integer, fraction = math.Modf(f)
	log.Printf("Modf(%v)=[%v,%v]", f, integer, fraction)

	f = 2
	integer, fraction = math.Modf(f)
	log.Printf("Modf(%v)=[%v,%v]", f, integer, fraction)

	f = -2.5
	integer, fraction = math.Modf(f)
	log.Printf("Modf(%v)=[%v,%v]", f, integer, fraction)

	f = 0.5
	integer, fraction = math.Modf(f)
	log.Printf("Modf(%v)=[%v,%v]", f, integer, fraction)
}

func example_Pow() {
	var a, b, res float64

	a = 2
	b = 10
	res = math.Pow(a, b)
	log.Printf("Pow(%v,%v)=[%v]", a, b, res)

	a = 1.5
	b = 2
	res = math.Pow(a, b)
	log.Printf("Pow(%v,%v)=[%v]", a, b, res)

	a = 3
	b = 0
	res = math.Pow(a, b)
	log.Printf("Pow(%v,%v)=[%v]", a, b, res)
}

func example_Signbit() {
	var f float64
	var res bool

	f = 4
	res = math.Signbit(f)
	log.Printf("Signbit(%v)=[%v]", f, res)

	f = -4
	res = math.Signbit(f)
	log.Printf("Signbit(%v)=[%v]", f, res)

	f = 0
	res = math.Signbit(f)
	log.Printf("Signbit(%v)=[%v]", f, res)

	f = -0
	res = math.Signbit(f)
	log.Printf("Signbit(%v)=[%v]", f, res)
}

func example_Min() {
	var a, b, res float64

	a = 2
	b = 3
	res = math.Min(a, b)
	log.Printf("Min(%v,%v)=[%v]", a, b, res)

	a = -2.1
	b = -3.3
	res = math.Min(a, b)
	log.Printf("Min(%v,%v)=[%v]", a, b, res)
}

func example_Max() {
	var a, b, res float64

	a = 2
	b = 3
	res = math.Max(a, b)
	log.Printf("Max(%v,%v)=[%v]", a, b, res)

	a = -2.1
	b = -3.3
	res = math.Max(a, b)
	log.Printf("Max(%v,%v)=[%v]", a, b, res)
}
