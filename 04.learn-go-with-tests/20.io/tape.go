package main

import (
	"io"
)

type tape struct {
	file io.ReadWriteSeeker
}

func (t *tape) Write(p []byte) (n int, err error) {
    //t.file.Truncate(0)
    t.file.Seek(0, 0)
    return t.file.Write(p)
}
