package sort

import (
	"fmt"
)

type Insertion struct {
	arr []int
}

func newInsertion(arr []int) *Insertion {
	insertion := &Insertion{
		arr: arr,
	}
	return insertion
}
func (m *Insertion) Sort() {
	len := len(m.arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if m.arr[j] > m.arr[i] {
				m.arr[j], m.arr[i] = m.arr[i], m.arr[j]
			}
		}
	}
}

func (m *Insertion) Array() []int {
	return m.arr
}

func (m *Insertion) String() string {
	result := ""
	for index, val := range m.arr {
		if index > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", val)
	}
	return result
}
