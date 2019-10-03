package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When 	time.Time
	What	string
}

/* 문자열(string)을 반환하는 하나의 메소드 Error 로 구성된 내장 인터페이스 타입 error 에서 나왔습니다.
type error interface {
	Error() string
}
*/
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work", 
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}