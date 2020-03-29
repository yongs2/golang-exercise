package main

import (
	"fmt"
)

type person struct {
	name   string
	gender string
	age    int
}

func main() {
	err := addPerson("Tina", "Female", 20)
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}

	err = addPerson("John", "Male")
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}

	err = addPerson("Wick", 2, 3)
	if err != nil {
		fmt.Println("PersonAdd Error: " + err.Error())
	}
}

func addPerson(args ...interface{}) error {
	if len(args) > 3 {
		return fmt.Errorf("Wront number of arguments passed")
	}
	p := &person{}
	//0 is name, 1 is gender, 2 is age
	for i, arg := range args {
		switch i {
		case 0: //name
			name, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Name is not passed as string")
			}
			p.name = name
		case 1: // gender
			gender, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Gender is not passed as string")
			}
			p.gender = gender
		case 2: // age
			age, ok := arg.(int)
			if !ok {
				return fmt.Errorf("Age is not passed as int")
			}
			p.age = age
		default:
			return fmt.Errorf("Wrong parametes passed")
		}
	}
	fmt.Printf("Person struct is %+v\n", p)
	return nil
}
