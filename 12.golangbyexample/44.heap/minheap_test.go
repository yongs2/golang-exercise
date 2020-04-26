package heap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinHeap(t *testing.T) {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	minHeap := newMinHeap(len(inputArray))

	t.Run("Insert MinHeap", func(t *testing.T) {
		for i := 0; i < len(inputArray); i++ {
			minHeap.insert(inputArray[i])
		}
		want := []int{2, 3, 5, 7, 6, 8}

		if reflect.DeepEqual(want, minHeap.heapArray) == false {
			t.Errorf("want[%+v], got[%+v]\n", want, minHeap.heapArray)
		}
	})

	t.Run("Remove MinHeap", func(t *testing.T) {
		minHeap.buildMinHeap()
		var result string
		for i := 0; i < len(inputArray); i++ {
			result += fmt.Sprintf("%d,", minHeap.remove())
		}
		t.Logf("Remove: %s\n", result)

		want := []int{}
		if reflect.DeepEqual(want, minHeap.heapArray) == false {
			t.Errorf("want[%+v], got[%+v]\n", want, minHeap.heapArray)
		}
	})
}
