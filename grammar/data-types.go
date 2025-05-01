package grammar

import (
	"fmt"
	"strings"
)

type point struct {
	x, y int
}

/*
StringFormat : 문자열 포맷팅
%v	인스턴스 출력	{1 2}
%+v	필드명 추가	{x:1 y:2}
%#v	코드 스니펫 출력	main.point{x:1, y:2}

%T	타입 출력	main.point

%t	불린 값	true
%d	10진수 출력	123
%b	바이너리 출력	1110
%c	정수에 해당하는 문자 출력	!
%x	16진수 인코딩 출력	1c8
%f	실수 출력	78.900000
%e	실수 출력	1.234000e+08
%E	실수 출력	1.234000E+08
%s	문자열 출력	"string"
%q	문자열에 쌍따옴표를 추가하여 출력	"\"string\""

%p	포인터 표현 출력	0xc0000160a0
%6d	우측 정렬로 숫자 출력
%6.2f	실수 정밀도 출력
%-6d	좌측 정렬로 숫자 출력
*/
func StringFormat() {
	p := point{1, 2}

	fmt.Printf("%%v 인스턴스 출력 %v\n", p)   // %v 인스턴스 출력 {1 2}
	fmt.Printf("%%+v 필드명 추가 %+v", p)    // %+v 필드명 추가 {x:1 y:2}
	fmt.Printf("%%T 타입 출력 %T\n", p)     // %T 타입 출력 main.point
	fmt.Printf("%%t 불린 값 %t\n", true)   // %t 불린 값 true
	fmt.Printf("%%d 10진수 출력 %d\n", 123) // %d 10진수 출력 123
	fmt.Printf("%%b 바이너리 출력 %b\n", 14)  // %b 바이너리 출력 1110
}

/*
StringsPackage : 문자열을 조작하고, 처리할 때는 strings 패키지를 이용
*/
func StringsPackage() {
	strs := []string{"a", "b", "c"}
	fmt.Println(strings.Join(strs, ":")) // 문자열 합치기 a:b:c

	str := "a.b.c"
	r := strings.Replace(str, ".", "_", -1) // 문자열 대치 a_b_c -1자리는 교체할 횟수 0보다 작으면 모두 교체
	fmt.Printf(r)
}
