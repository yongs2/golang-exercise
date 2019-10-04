package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 어떤 식으로든 스트림을 수정하여 다른 io.Reader 를 감싸는 io.Reader 는 흔한 패턴입니다.
// ROT13 치환 암호화를 모든 알파벳 문자에 적용함으로써 스트림을 수정
type rot13Reader struct {
	r io.Reader
}

func lookup(b byte) (byte) {
	//Input"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	str := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	var idx int = 0
	var startIdx int = len("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	if b >= 'A' && b <= 'Z' {
		idx = int(b - 'A')
	} else if b >= 'a' && b <= 'z' {
		idx = startIdx + int(b - 'a')
	} else {
		return b
	}
	//fmt.Println("Org:", string(b), ", Dest:", string(str[idx]))
	return str[idx]
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	//fmt.Println(">>Read=", n)
	for i := 0; i<n; i++ {
		b[i] = lookup(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//s := strings.NewReader("The Quick Brown Fox Jumps Over The Lazy Dog.")
	//s := strings.NewReader("Gur Dhvpx Oebja Sbk Whzcf Bire Gur Ynml Qbt.")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println("")
}
