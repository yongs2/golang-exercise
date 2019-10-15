package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40}) // Pointer
	fmt.Println(NewPerson("Jon"))              // Pointer

	// j is pointer
	j := NewPerson("John")
	fmt.Println(j, j.name, j.age)
	j.name = "John2"
	j.age = j.age + 1
	fmt.Println(j, j.name, j.age)

	// s is struct
	s := person{name: "Sean", age: 50}
	fmt.Println(s, s.name, s.age)
	s.name = "Sean2"
	s.age = 49
	fmt.Println(s, s.name, s.age)

	// sp is pointer
	sp := &s
	fmt.Println(sp, sp.name, sp.age)

	sp.age = 51
	fmt.Println(sp, sp.name, sp.age)
}
