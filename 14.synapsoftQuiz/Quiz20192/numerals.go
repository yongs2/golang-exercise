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

var koreanNum = []string{"", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구", "십"}
var koreanVal = []string{"", "십", "백", "천", "만", "십만", "백만", "천만", "억", "십억", "백억", "천억", "조", "십조", "백조", "천조"}

func ConvertKoreanNumerals(inputNum string) (korean string) {
	korean = ""

	digitStr := strings.Split(inputNum, "원")
	if len(digitStr) < 2 {
		return ""
	}
	log.Printf("digitStr[%v]=[%v]\n", len(digitStr), digitStr[0])
	idx := 0
	numIdx := 0
	bNotOne := false
	for i := len(digitStr[0]) - 1; i >= 0; i-- {
		if digitStr[0][i] == ',' {
			continue
		}
		numIdx = int(digitStr[0][i] - '0')

		if numIdx == 0 {

		} else {
			if numIdx == 1 { // 일
				if idx > 0 && idx < 5 {
					bNotOne = true
				}
			}

			if (idxs+1)%4 == 0 {
				korean += " "
			}
			if bNotOne == true {
				korean = koreanVal[idx] + korean
			} else {
				korean = koreanNum[numIdx] + koreanVal[idx] + korean
			}

			//log.Printf("korean[%v], numIdx[%v],val[%v]\n", korean, numIdx, idx)
		}

		idx++
		//log.Printf("digit[%v]=[%c]\n", i, digitStr[0][i])
	}

	korean = korean + "원"
	return
}
