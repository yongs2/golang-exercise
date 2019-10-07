package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func (t *tree.Tree) {
        if (t == nil) {
            return
        }
        walker(t.Left)
        ch <- t.Value
        walker(t.Right)
    }
    walker(t)
    close(ch)	// 채널을 닫아야 하므로, walker함수로 내부를 다시 정의해야 함
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// 결과를 문자열로 변환
	s1 := ""
	for i := range ch1 {
		s1 += fmt.Sprintf("%d ", i)
	}
	fmt.Println("s1=", s1)

	// 결과를 문자열로 변환
	s2 := ""
	for i := range ch2 {
		s2 += fmt.Sprintf("%d ", i)
	}
	fmt.Println("s2=", s2)

	// 문자열 비교
    return (s1 == s2)
}

func main() {
	var max int = 10

	ch := make(chan int, max)
	t := tree.New(max)

	fmt.Println(t)

	go Walk(t, ch)
	for i := range ch {	// 반복문은 채널이 닫힐 때까지 계속해서 값을 받습니다.
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(10), tree.New(10)))
}