package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	inputArray := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	Bubble := newBubble(inputArray)
	Bubble.Sort()

	got := Bubble.Array()
	want := []int{-3, -1, 1, 2, 3, 4, 5, 7, 8}
	if reflect.DeepEqual(want, got) == false {
		t.Errorf("want[%+v], got[%+v]\n", want, got)
	}
}

// go test -bench=. -benchmem
func BenchmarkBubbleSort(b *testing.B) {
	inputArray := generateRandomNumber(1000)
	for i := 0; i < b.N; i++ {
		Bubble := newBubble(inputArray)
		Bubble.Sort()
	}
}
