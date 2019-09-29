package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibaonacci() func() int {
	index := 0
	fn0 := 0
	fn1 := 0
	fn2 := 0
	
	return func() int {
		if index == 0 {
			fn0 = 0;
		} else if index == 1 {
			fn0 = 1;
		} else if index > 1 {
			fn2 = fn1;
			fn1 = fn0;
			fn0 = fn1 + fn2
		}
		//fmt.Println("Index=", index, "fn1=", fn1, "fn2=", fn2)
		index += 1
		return fn0
	}
}

func main() {
	f := fibaonacci()
	for i := 0; i < 13; i++ {
		fmt.Println(f())
	}
}