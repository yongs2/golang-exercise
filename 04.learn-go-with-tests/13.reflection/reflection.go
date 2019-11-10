package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
			fmt.Println("walk:", i, ":", val.Field(i), val.Field(i).Kind())
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
			fmt.Println("walk:", i, ":", val.Index(i), val.Index(i).Kind())
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	fmt.Println("walk:getValue:", val)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
