package calc

import (
	"testing"
)

// go test -v ./... -coverprofile=coverage.out
// go tool cover -html=coverage.out -o coverage.html
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
		input int
		want  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{20, "XX"},
		{30, "XXX"},
		{39, "XXXIX"},
	}

	for _, tc := range romaniTests {
		got := ConvertRomani(tc.input)
		if got != tc.want {
			t.Errorf("got %v want %v", got, tc.want)
		}
	}
}

func TestParseInput(t *testing.T) {
	parseInputTests := []struct {
		input        string
		wantLeft     string
		wantOperator string
		wantRight    string
		wantError    error
	}{
		{"", "", "", "", ErrNoInput},
		{"II ^ II", "", "", "", ErrInvalidOperator},
		{"I + I", "I", "+", "I", nil},
		{"XX - IV", "XX", "-", "IV", nil},
		{"III * VI", "III", "*", "VI", nil},
		{"X / III", "X", "/", "III", nil},
		{"III / X", "III", "/", "X", nil},
		{"X - XXX", "X", "-", "XXX", nil},
		{"XX + XX", "XX", "+", "XX", nil},
	}

	for _, tc := range parseInputTests {
		gotLeft, gotOperator, gotRight, gotError := ParseInput(tc.input)
		if gotLeft != tc.wantLeft {
			t.Errorf("got %q want %q", gotLeft, tc.wantLeft)
		}
		if gotOperator != tc.wantOperator {
			t.Errorf("got %q want %q", gotOperator, tc.wantOperator)
		}
		if gotRight != tc.wantRight {
			t.Errorf("got %q want %q", gotRight, tc.wantRight)
		}
		if gotError != tc.wantError {
			t.Errorf("got %q want %q", gotError, tc.wantError)
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
