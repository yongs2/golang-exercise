package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	fmt.Println("walk:", val)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Println("walk:", i, ":", field, field.Kind())
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
