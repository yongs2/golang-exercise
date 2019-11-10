package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	fmt.Println("walk:", val)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Println("walk:", i, ":", field, field.Kind())
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
