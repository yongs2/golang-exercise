package calc

import (
	"testing"
)

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
		{"X – XXX", "작은 수에서 큰 수를 뺄 수 없습니다."},
		{"XX + XX", "범위를 벗어났습니다."},
	}

	for _, tc := range calcTests {
		got := Calc(tc.input)
		if got != tc.want {
			t.Errorf("got %q want %q", got, tc.want)
		}
	}
}
