package reflection

import (
	"reflect"
	"fmt"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	fmt.Println("walk:", val)
	
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Println("walk:", i, ":", field)
		fn(field.String())
	}
}
