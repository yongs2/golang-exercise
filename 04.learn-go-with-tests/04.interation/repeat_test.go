package interation

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrect := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if expected != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat 5", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrect(t, repeated, expected)
	})
	
	t.Run("repeat 0", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "a"
		assertCorrect(t, repeated, expected)
	})

	t.Run("repeat -1", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "a"
		assertCorrect(t, repeated, expected)
	})
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
