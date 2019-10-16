package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// start with a string representation of our JSON data
var input = `
{
	"created_at" : "Thu May 31 00:00:01 +0000 2012"
}
`

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func main() {
	var val map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}
	fmt.Println("V:", val)
	for k, v := range val {
		fmt.Println("F:", k, reflect.TypeOf(k), v, reflect.TypeOf(v))
	}
	fmt.Println("1:", Timestamp(val["created_at"]))
	fmt.Println("2:", time.Time(val["created_at"]))
}
