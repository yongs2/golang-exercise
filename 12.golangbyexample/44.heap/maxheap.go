package heap

import (
	"fmt"
)

type MaxHeap struct {
	heapArray []int
	size      int
	maxSize   int
}

func newMaxHeap(maxSize int) *MaxHeap {
	MaxHeap := &MaxHeap{
		heapArray: []int{},
		size:      0,
		maxSize:   maxSize,
	}
	return MaxHeap
}

func (m *MaxHeap) leaf(index int) bool {
	if index >= (m.size/2) && index <= m.size {
		return true
	}
	return false
}

func (m *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *MaxHeap) leftChild(index int) int {
	return 2*index + 1
}

func (m *MaxHeap) rightChild(index int) int {
	return 2*index + 2
}

func (m *MaxHeap) insert(item int) error {
	if m.size >= m.maxSize {
		return fmt.Errorf("Heap is full")
	}
	m.heapArray = append(m.heapArray, item)
	m.size++
	m.upHeapify(m.size - 1)
	return nil
}

func (m *MaxHeap) swap(first, second int) {
	temp := m.heapArray[first]
	m.heapArray[first] = m.heapArray[second]
	m.heapArray[second] = temp
}

func (m *MaxHeap) upHeapify(index int) {
	for m.heapArray[index] > m.heapArray[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *MaxHeap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	largest := current
	leftChildIndex := m.leftChild(current)
	rightChildIndex := m.rightChild(current)
	if leftChildIndex < m.size && m.heapArray[leftChildIndex] > m.heapArray[largest] {
		largest = leftChildIndex
	}
	if rightChildIndex < m.size && m.heapArray[rightChildIndex] > m.heapArray[largest] {
		largest = rightChildIndex
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
	return
}

func (m *MaxHeap) buildMaxHeap() {
	for index := ((m.size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *MaxHeap) remove() int {
	top := m.heapArray[0]
	m.heapArray[0] = m.heapArray[m.size-1]
	m.heapArray = m.heapArray[:(m.size)-1]
	m.size--
	m.downHeapify(0)
	return top
}
