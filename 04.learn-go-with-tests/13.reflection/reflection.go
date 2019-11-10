package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
		fmt.Println("walk:", i, ":", getField(i), getField(i).Kind())
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
