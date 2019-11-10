package reflection

import (
	"reflect"
	"fmt"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	fmt.Println("walk:", val)
	field := val.Field(0)
	fmt.Println("walk:", field)
	fn(field.String())
}
