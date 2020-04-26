package heap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	inputArray := []int{6, 5, 3, 7, 2, 8}
	maxHeap := newMaxHeap(len(inputArray))

	t.Run("Insert MaxHeap", func(t *testing.T) {
		for i := 0; i < len(inputArray); i++ {
			maxHeap.insert(inputArray[i])
		}
		want := []int{8, 6, 7, 5, 2, 3}

		if reflect.DeepEqual(want, maxHeap.heapArray) == false {
			t.Errorf("want[%+v], got[%+v]\n", want, maxHeap.heapArray)
		}
	})

	t.Run("Remove MaxHeap", func(t *testing.T) {
		maxHeap.buildMaxHeap()
		var result string
		for i := 0; i < len(inputArray); i++ {
			result += fmt.Sprintf("%d,", maxHeap.remove())
		}
		t.Logf("Remove: %s\n", result)

		want := []int{}
		if reflect.DeepEqual(want, maxHeap.heapArray) == false {
			t.Errorf("want[%+v], got[%+v]\n", want, maxHeap.heapArray)
		}
	})
}
