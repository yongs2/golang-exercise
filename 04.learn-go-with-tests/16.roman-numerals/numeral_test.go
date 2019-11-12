package numeral

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("1 gets convert to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	
	t.Run("2 gets convert to II", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
