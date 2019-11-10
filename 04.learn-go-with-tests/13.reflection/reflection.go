package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
			fmt.Println("walk:", i, ":", val.Field(i), val.Field(i).Kind())
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
			fmt.Println("walk:", i, ":", val.Index(i), val.Index(i).Kind())
		}
	case reflect.Map:
		i := 0
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
			fmt.Println("walk:", i, ":", val.MapIndex(key), val.MapIndex(key).Kind())
			i += 1
		}
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
