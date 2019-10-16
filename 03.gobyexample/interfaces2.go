package main

// Ref : https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Wooof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgammar struct {
}

func (j JavaProgammar) Speak() string {
	return "Design Patterns!"
}

func main() {
	//animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgammar{}}
	animals := []Animal{new(Dog), new(Cat), new(Llama), new(JavaProgammar)}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
