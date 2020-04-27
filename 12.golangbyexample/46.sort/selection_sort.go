package sort

import (
	"fmt"
)

type Selection struct {
	arr []int
}

func newSelection(arr []int) *Selection {
	selection := &Selection{
		arr: arr,
	}
	return selection
}
func (m *Selection) Sort() {
	len := len(m.arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if m.arr[j] < m.arr[minIndex] {
				m.arr[j], m.arr[minIndex] = m.arr[minIndex], m.arr[j]
			}
		}
	}
}

func (m *Selection) Array() []int {
	return m.arr
}

func (m *Selection) String() string {
	result := ""
	for index, val := range m.arr {
		if index > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", val)
	}
	return result
}
