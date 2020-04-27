package sort

import (
	"fmt"
)

type Bubble struct {
	arr []int
}

func newBubble(arr []int) *Bubble {
	selection := &Bubble{
		arr: arr,
	}
	return selection
}
func (m *Bubble) Sort() {
	len := len(m.arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if m.arr[j] > m.arr[j+1] {
				m.arr[j], m.arr[j+1] = m.arr[j+1], m.arr[j]
			}
		}
	}
}

func (m *Bubble) Array() []int {
	return m.arr
}

func (m *Bubble) String() string {
	result := ""
	for index, val := range m.arr {
		if index > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", val)
	}
	return result
}
