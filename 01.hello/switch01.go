package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// case의 코드 실행을 마치면 알아서 break를 한다는 점
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X")
	case "linux" : 
		fmt.Println("Linux")
	default :
		// freebsd, openbsd, plan9, windows
		fmt.Printf("%s.", os)
	}
}