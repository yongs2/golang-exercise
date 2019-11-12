package numeral

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets convert to I", 1, "I"},
		{"2 gets convert to II", 2, "II"},
		{"3 gets convert to III", 3, "III"},
		{"4 gets convert to IV", 4, "IV"},
		{"5 gets convert to V", 5, "V"},
		{"6 gets convert to VI", 6, "VI"},
		{"7 gets convert to VII", 7, "VII"},
		{"8 gets convert to VIII", 8, "VIII"},
		{"9 gets convert to IX", 9, "IX"},
		{"10 gets convert to X", 10, "X"},
		{"14 gets convert to XIV", 14, "XIV"},
		{"18 gets convert to XVIII", 18, "XVIII"},
		{"20 gets convert to XX", 20, "XX"},
		{"39 gets convert to XXXIX", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
