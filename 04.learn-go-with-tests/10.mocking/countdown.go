package main

import (
	"fmt"
	"bytes"
)

func Countdown(out *bytes.Buffer) {
	fmt.Fprintf(out, "3")
}
