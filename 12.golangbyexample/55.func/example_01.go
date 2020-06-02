package main

import (
	"fmt"
	"log"
)

func main() {
	example_immediately()
	example_Closures()
	example_ValueOutside()
	example_AccessVariable()
	example_Func()
	example_ReturnFunc()
	example_VariadicFuncSameType()
	example_VariadicFuncDifferentType()
	example_Person()
}

func example_immediately() {
	squareOf2 := func() int {
		return 2 * 2
	}()
	log.Println("squareOf2:", squareOf2)
}

func example_Closures() {
	modulus := getModulus()
	log.Printf("modulus=%v", modulus(-1))
	log.Printf("modulus=%v", modulus(2))
	log.Printf("modulus=%v", modulus(-5))
}

func getModulus() func(int) int {
	count := 0
	return func(x int) int {
		count = count + 1
		log.Printf("modulus function called %d times", count)
		if x < 0 {
			x = x * -1
		}
		return x
	}
}

func example_ValueOutside() {
	valueOutside := "somevalue"
	func() {
		log.Printf(valueOutside)
	}()
}

func example_AccessVariable() {
	count := 0
	for i := 1; i <= 5; i++ {
		func() {
			count++
			log.Printf("i=%v, count=%v", i, count)
		}()
	}
}

func example_Func() {
	print(area, 2, 4)
	print(sum, 2, 4)
}

func print(f func(int, int) int, a int, b int) {
	log.Printf("[%V](%v, %v)=%v", f, a, b, f(a, b))
}

func area(a int, b int) int {
	return a * b
}

func sum(a int, b int) int {
	return a + b
}

func example_ReturnFunc() {
	areaF := getAreaFunc()
	res := areaF(2, 4)
	log.Printf("areaF(%v, %v)=%v", 2, 4, res)
}

func getAreaFunc() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}

func example_VariadicFuncSameType() {
	log.Printf("add(1,2)=%v", add(1, 2))
	log.Printf("add(1,2,3)=%v", add(1, 2, 3))
	log.Printf("add(1,2,3,4)=%v", add(1, 2, 3, 4))
}

func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func example_VariadicFuncDifferentType() {
	handle(1, "abc")
	handle("abc", "xyz", 3)
	handle(1, 2, 3, 4)
}

func handle(params ...interface{}) {
	log.Printf("Handle func called with parameters:\n")
	for _, param := range params {
		log.Printf("%v\n", param)
	}
}

func example_Person() {
	var err error

	if err = addPerson("Tina", "Female", 20); err != nil {
		log.Printf("PersonAdd Error[%v]", err)
	}

	if err = addPerson("John", "Male"); err != nil {
		log.Printf("PersonAdd Error[%v]", err)
	}

	if err = addPerson("Wick", 2, 3); err != nil {
		log.Printf("PersonAdd Error[%v]", err)
	}
}

type person struct {
	name   string
	gender string
	age    int
}

func addPerson(args ...interface{}) error {
	if len(args) > 3 {
		return fmt.Errorf("Wront number of arguments passed")
	}

	p := &person{}
	for i, arg := range args {
		switch i {
		case 0:
			name, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Name is not passed as string")
			}
			p.name = name
		case 1:
			gender, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Gender is not passed as string")
			}
			p.gender = gender
		case 2:
			age, ok := arg.(int)
			if !ok {
				return fmt.Errorf("Age is not passed as int")
			}
			p.age = age
		default:
			return fmt.Errorf("Wrong parametes passed")
		}
	}
	log.Printf("Person struct is %+v\n", p)
	return nil
}
