package model

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	fmt.Println(p.Name)
	fmt.Println(p.age)

	person2 := GetPerson()
	fmt.Println(person2)
	companyName := getCompanyName()
	fmt.Println(companyName)

	fmt.Println(companyLocation)
	fmt.Println(CompanyName)
}
