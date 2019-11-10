package reflection

import (
	"testing"
	"fmt"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	fmt.Println("got=", got[0])
	fmt.Println("expected=", expected)
	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}
}
