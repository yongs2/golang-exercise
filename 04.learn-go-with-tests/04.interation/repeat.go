package interation

const repeatDefaultCount = 1

func Repeat(character string, repeatCount int) string {
	var repeated string

	if repeatCount <= 0 {
		repeatCount = repeatDefaultCount
	}
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
