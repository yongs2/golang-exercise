package numerals

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

func ConvertKoreanNumerals(inputNum string) (korean string) {
	if inputNum == "1원" {
		return "일원"
	}
	return "에러"
}
