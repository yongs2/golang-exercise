package main

import (
	"fmt"
	"os"
)

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

type T struct {
	a int
	b float64
	c string
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	fmt.Println("unknown time zone:", tz)
	return 0
}

func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

func main() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	fmt.Printf("%v\n", timeZone) // or just fmt.Println(timeZone)

	offset1 := timeZone["EST"]
	fmt.Println("EST=", offset1, offset("EST"))

	t := &T{7, -2.35, "abc\tdef"}

	// 포맷별로 출력 결과 비교
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)
	fmt.Printf("%T\n", timeZone)

	// 커스텀 타입
	fmt.Printf("%v\n", t)
}
