package main

import (
    "golang.org/x/tour/wc"
    //"fmt"
    "strings"
)

// 이 함수는 s라는 문자열 내에서 각각의 "단어"의 등장 횟수를 나타내는 맵을 반환해야 합니다.
func WordCount(s string) map[string]int {
	//fmt.Println("input=", s)
	
	// https://golang.org/pkg/strings/#Fields
    var a = strings.Fields(s)
    //fmt.Println("Field=", a, len(a))
    
    m := make(map[string] int)
    
    var i, j, nSame int
    
    for i=0; i < len(a); i++ {
        //fmt.Println("i=", i, "str=", a[i])
        nSame = 0
        for j=0; j < len(a); j++ {
            if a[i] == a[j] {
                nSame += 1
            }
        }
        m[a[i]] = nSame
    }
    
    return m
}

func main() {
    wc.Test(WordCount)
}