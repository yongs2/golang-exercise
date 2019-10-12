package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf	// For Debugging; delete when done.
var _ io.Reader	// For Debuggin; delete when done.

func main() {
	fd, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: use fd
	_ = fd 	// For Debugging; delete when done.
}
