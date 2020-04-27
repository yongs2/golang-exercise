package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	inputArray := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	Insertion := newInsertion(inputArray)
	Insertion.Sort()

	got := Insertion.Array()
	want := []int{-3, -1, 1, 2, 3, 4, 5, 7, 8}
	if reflect.DeepEqual(want, got) == false {
		t.Errorf("want[%+v], got[%+v]\n", want, got)
	}
}

// go test -bench=. -benchmem
func BenchmarkInsertionSort(b *testing.B) {
	inputArray := generateRandomNumber(1000)
	for i := 0; i < b.N; i++ {
		Insertion := newInsertion(inputArray)
		Insertion.Sort()
	}
}
