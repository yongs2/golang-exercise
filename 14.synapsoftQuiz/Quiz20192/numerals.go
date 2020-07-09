package numerals

import (
	"log"
	"strings"
)

/* https://www.synapsoft.co.kr/blog/12346
 * korean numerals
 * 80,270원 처럼 금액의 표기는 천 단위로 콤마를 찍지만
 * 실제로 읽을 때는 ‘팔만 이백칠십원’처럼 만 단위 변환
 * 1억원, 1조원을 일억 원, 일조 원
 * 1만원, 1천원, 1백원의 경우는 만원, 천원, 백원, 십원
 *
 * 입력 금액의 범위는 1원에서 100조원
 * 모든 금액은 천 단위 구분자인 콤마가 표시돼있고 금액단위인 원
 */

var koreanNum = []string{"", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구"}
var koreanVal = []string{"", "십", "백", "천"}
var koreanDan = []string{"", "만", "억", "조"}

func ConvertKoreanNumerals(inputNum string) (korean string) {
	korean = ""

	digitStr := strings.Split(inputNum, "원")
	if len(digitStr) < 2 {
		return ""
	}
	//log.Printf("digitStr[%v]=[%v]\n", len(digitStr), digitStr[0])
	idx := 0
	numIdx := 0
	valIdx := 0
	danIdx := 0
	insertDanIdx := 0
	bNotOne := false
	for i := len(digitStr[0]) - 1; i >= 0; i-- {
		bNotOne = false
		if digitStr[0][i] == ',' {
			continue
		}
		numIdx = int(digitStr[0][i] - '0')
		if numIdx == 0 {

		} else {
			if numIdx == 1 { // 일
				if idx > 0 && idx < 4 {
					bNotOne = true
				}
			}
			valIdx = idx % 4
			danIdx = idx / 4

			if idx != 0 && valIdx == 0 {
				if len(korean) > 0 {
					korean = " " + korean
				}
			}
			//log.Printf("korean[%v], numIdx[%v],dan[%v],val[%v],idx[%v]\n", korean, numIdx, danIdx, valIdx, idx)
			if danIdx != insertDanIdx {
				korean = koreanDan[danIdx] + korean
			}
			korean = koreanVal[valIdx] + korean
			if bNotOne == false {
				korean = koreanNum[numIdx] + korean
			}
			insertDanIdx = danIdx
		}
		idx++
		//log.Printf("korean[%v], bNotOne[%v]\n", korean, bNotOne)
	}

	if valIdx >= 1 && valIdx <= 4 || danIdx == 1 { // "일만" 인 경우는 일 제거
		korean = strings.TrimLeft(korean, "일")
	}

	korean = korean + "원"
	log.Printf("korean[%v]\n", korean)
	return
}
