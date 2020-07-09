package numerals

import (
	"testing"
)

// go test -v ./... -coverprofile=coverage.out
// go tool cover -html=coverage.out -o coverage.html
func TestConvertKoreanNumerals(t *testing.T) {
	numeralsTests := []struct {
		input string
		want  string
	}{
		{"1원", "일원"},
		{"1,000원", "천원"},
		{"10,000원", "만원"},
		{"1,000,000원", "백만원"},
		{"100,000,000원", "일억원"},
		{"80,270원", "팔만 이백칠십원"},
		{"111,111원", "십일만 천백십일원"},
		{"1,234,567,890원", "십이억 삼천사백오십육만 칠천팔백구십원"},
		{"100,000,000,000,000원", "백조원"},
	}

	for _, tc := range numeralsTests {
		got := ConvertKoreanNumerals(tc.input)
		if got != tc.want {
			t.Errorf("got %v want %v", got, tc.want)
		}
	}
}
