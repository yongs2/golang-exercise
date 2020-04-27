package sort

import (
	"fmt"
)

// Time Complexity of Heap Sort is  O(nLogn).
type minHeap struct {
	arr []int
}

func newMinHeap(arr []int) *minHeap {
	minHeap := &minHeap{
		arr: arr,
	}
	return minHeap
}

func (m *minHeap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (m *minHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (m *minHeap) swap(first, second int) {
	temp := m.arr[first]
	m.arr[first] = m.arr[second]
	m.arr[second] = temp
}

func (m *minHeap) leaf(index int, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}
	return false
}

func (m *minHeap) downHeapify(current int, size int) {
	if m.leaf(current, size) {
		return
	}
	smallest := current
	leftChildIndex := m.leftChildIndex(current)
	rightChildIndex := m.rightChildIndex(current)
	if leftChildIndex < size && m.arr[leftChildIndex] < m.arr[smallest] {
		smallest = leftChildIndex
	}
	if rightChildIndex < size && m.arr[rightChildIndex] < m.arr[smallest] {
		smallest = rightChildIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest, size)
	}
	return
}

func (m *minHeap) buildMinHeap(size int) {
	for index := ((size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index, size)
	}
}

func (m *minHeap) sort(size int) {
	m.buildMinHeap(size)
	for i := size - 1; i > 0; i-- {
		m.swap(0, i)
		m.downHeapify(0, i)
	}
}

func (m *minHeap) Array() []int {
	return m.arr
}

func (m *minHeap) String() string {
	result := ""
	for index, val := range m.arr {
		if index > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", val)
	}
	return result
}
