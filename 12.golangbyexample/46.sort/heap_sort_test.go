package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	inputArray := []int{6, 5, 3, 7, 2, 8, -1}
	minHeap := newMinHeap(inputArray)
	minHeap.sort(len(inputArray))

	got := minHeap.Array()
	want := []int{8, 7, 6, 5, 3, 2, -1}
	if reflect.DeepEqual(want, got) == false {
		t.Errorf("want[%+v], got[%+v]\n", want, got)
	}
}

// go test -bench=. -benchmem
func BenchmarkHeapSort(b *testing.B) {
	inputArray := generateRandomNumber(1000)
	for i := 0; i < b.N; i++ {
		minHeap := newMinHeap(inputArray)
		minHeap.sort(len(inputArray))
	}
}
