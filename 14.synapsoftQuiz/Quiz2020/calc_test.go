package calc

import (
	"testing"
)

func TestParseRomani(t *testing.T) {
	romaniTests := []struct {
		input string
		want  int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"VI", 6},
		{"VII", 7},
		{"VIII", 8},
		{"IX", 9},
		{"X", 10},
		{"XX", 20},
		{"XXX", 30},
		{"XXXIX", 39},
	}

	for _, tc := range romaniTests {
		got, _ := ParseRomani(tc.input)
		if got != tc.want {
			t.Errorf("got %v want %v", got, tc.want)
		}
	}
}

func TestConvertRomani(t *testing.T) {
	romaniTests := []struct {
		want  string
		input int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"VI", 6},
		{"VII", 7},
		{"VIII", 8},
		{"IX", 9},
		{"X", 10},
		{"XX", 20},
		{"XXX", 30},
		{"XXXIX", 39},
	}

	for _, tc := range romaniTests {
		got := ConvertRomani(tc.input)
		if got != tc.want {
			t.Errorf("got %v want %v", got, tc.want)
		}
	}
}

func TestCalc(t *testing.T) {
	calcTests := []struct {
		input string
		want  string
	}{
		{"I + I", "II"},
		{"XX - IV", "XVI"},
		{"III * VI", "XVIII"},
		{"X / III", "몫 III, 나머지 I"},
		{"III / X", "작은 수를 큰 수로 나눌 수 없습니다."},
		{"X - XXX", "작은 수에서 큰 수를 뺄 수 없습니다."},
		{"XX + XX", "범위를 벗어났습니다."},
	}

	for _, tc := range calcTests {
		got := Calc(tc.input)
		if got != tc.want {
			t.Errorf("got %q want %q", got, tc.want)
		}
	}
}
