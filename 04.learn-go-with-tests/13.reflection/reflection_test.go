package reflection

import (
	"testing"
	"reflect"
	"fmt"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"Struct with one string field",
			struct {
				Name string
			} { "Chris" },
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			} { "Chris", "London" },
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			fmt.Println("test:Input:", test.Input, "test:Expected:", test.ExpectedCalls)
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			fmt.Println("test:got:", got)
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
