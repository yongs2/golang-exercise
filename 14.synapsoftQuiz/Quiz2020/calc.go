package calc

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

/*
 * 로마 숫자 계산기
 *
 * 조건
 * 사칙연산이 가능
 * 입력값과 결과값의 범위는 I ~ XXXIX (1 ~ 39)
 * 결과값이 0이하 이거나 39(XXXIX)보다 큰 경우 “범위를 벗어났습니다.”라고 표시
 * 작은 수에서 큰 수를 빼는 경우 “작은 수에서 큰수를 뺄 수 없습니다.”라고 표시
 * 작은 수를 큰 수로 나누는 경우 “작은 수를 큰 수로 나눌 수 없습니다.”라고 표시
 * 나누기(/)의 결과 값은 몫과 나머지로 표시
 * 입력은 파일이나 콘솔에서 받아도 되고 프로그램 내부에 하드코딩
 *
 */

// MaxValue : 최대 값
const MaxValue int = 39

// 에러 정의
var ErrNoInput = errors.New("입력값이 없습니다.")
var ErrInvalidOperator = errors.New("연산자를 잘못 입력하였습니다.")
var ErrInvalidRomani = errors.New("로마숫자 형식이 잘못되었습니다.")

// Calc : 로마 숫자 계산기
func Calc(input string) (result string) {
	var err error

	left, operator, right, err := ParseInput(input)
	if err != nil {
		return err.Error()
	}

	var a, b, c, d int
	if a, err = ParseRomani(left); err != nil {
		return "잘못 입력하였습니다.2"
	}
	if b, err = ParseRomani(right); err != nil {
		return "잘못 입력하였습니다.3"
	}

	// 사칙연산
	if operator == "+" {
		c = a + b
	} else if operator == "-" {
		if a < b {
			return "작은 수에서 큰 수를 뺄 수 없습니다."
		}
		c = a - b
	} else if operator == "*" {
		c = a * b
	} else if operator == "/" {
		if a < b {
			return "작은 수를 큰 수로 나눌 수 없습니다."
		}
		c = a / b
		d = a % b
	} else {
		return "잘못 입력하였습니다."
	}
	log.Printf("%v %v %v = %v (%d)", a, operator, b, c, d)

	if c > MaxValue {
		return "범위를 벗어났습니다."
	}
	if operator == "/" {
		return fmt.Sprintf("몫 %v, 나머지 %v", ConvertRomani(c), ConvertRomani(d))
	}
	return ConvertRomani(c)
}

// ParseInput : 로마 숫자 계산식을 사칙연산자를 기준으로 분리
func ParseInput(input string) (left, operator, right string, err error) {
	if len(input) <= 0 {
		return "", "", "", ErrNoInput
	}

	// 사칙연산 을 기준으로 값을 구분
	nOffset := strings.IndexAny(input, "+-*/")
	if nOffset < 0 {
		return "", "", "", ErrInvalidOperator
	}

	operator = string(input[nOffset])
	left = strings.TrimSpace(input[:nOffset])
	right = strings.TrimSpace(input[nOffset+1:])
	//log.Printf("op[%v], left[%v], right[%v]", operator, left, right)
	return left, operator, right, nil
}

// ParseRomani : 로마 숫자를 integer 로 변환
func ParseRomani(input string) (result int, err error) {
	var prevChar byte
	for _, char := range input {
		//log.Printf("index=%d, char=[%c]", index, char)
		if char == 'I' {
			result++
			prevChar = 'I'
		} else if char == 'V' {
			if prevChar == 'I' {
				result -= 1 // 이전에 I 에 대해 1로 추가한 부분을 제거
				result += 4 // IV 인 경우 4로 계산
				prevChar = ' '
			} else {
				result += 5 // V로 시작하면 5로 계산
				prevChar = 'V'
			}
		} else if char == 'X' {
			if prevChar == 'I' {
				result -= 1
				result += 9
				prevChar = ' '
			} else {
				result += 10
				prevChar = 'X'
			}
		}
	}
	//log.Printf("result=[%v]", result)
	if result < 0 {
		return -1, ErrInvalidRomani
	}
	return result, nil
}

// Romani : integer를 로마 숫자로 변환하기 위한 문자열 배열
var Romani = [11]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

// ConvertRomani : integer를 로마 숫자 문자열로 변환
func ConvertRomani(input int) (output string) {
	division := input / 10
	extra := input % 10

	for i := 0; i < division; i++ {
		output += Romani[10]
	}
	output += Romani[extra]
	//log.Printf("ConvertRomani(%v)=[%v]\n", input, output)
	return output
}
