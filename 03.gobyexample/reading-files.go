package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Test
// echo "hello" > /tmp/dat; echo "go" >> /tmp/dat
func main() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("/tmp/dat")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	var nPos int64 = 6	// hello\ngo\n
	var nMaxRead int = 2

	o2, err := f.Seek(nPos, 0)
	check(err)
	b2 := make([]byte, nMaxRead)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(nPos, 0)
	check(err)
	b3 := make([]byte, nMaxRead)
	n3, err := io.ReadAtLeast(f, b3, len(b3))
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// rewind
	_, err = f.Seek(0, 0)
	check(err)

	nRead := 5
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(nRead)
	check(err)
	fmt.Printf("%d bytes: %s\n", nRead, string(b4))

	f.Close()
}
