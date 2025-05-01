package grammar

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
StringToInt : 문자열 정수형으로 변환
	함수				반환 	타입	진수(base)	bitSize 지정 가능?			주요 용도
strconv.Atoi		int		항상 10진수			❌				단순히 문자열을 int로 변환할 때
strconv.ParseInt	int64	원하는 진수 지정 가능	✅				다양한 진수, 크기 변환할 때
*/
func StringToInt() {
	strInt := "100"
	// 문자열 → 정수 변환: strconv.Atoi를 사용
	i, err := strconv.Atoi(strInt)
	fmt.Println(i, err, reflect.TypeOf(i)) // 100 <nil> int

	strInt = "987654321"
	// 8비트 크기로 해석 (최대값 127). 오버플로 되므로 최대값인 127이 나옴.
	i8, err := strconv.ParseInt(strInt, 0, 8)
	fmt.Println(i8, err, reflect.TypeOf(i8)) // 127 <nil> int64
}
