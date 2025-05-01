package grammar

import "fmt"

/*
var 키워드를 통해 변수 선언 / 변수명 := 값 을 통해 선언
변수를 선언하고 초기화하지 않으면 기본값 (int 0, String "", bool false)
*/
func Varivales() {
	var i1 int = 10
	var s1 string = "string"

	fmt.Println(i1, s1)

	// 타입 생략 가능
	var i2 = 10
	var s2 = "string"

	fmt.Println(i2, s2)

	// := 를 통한 변수 선언
	i3 := 10
	s3 := "string"

	fmt.Println(i3, s3)

	// 다수 변수 동시 선언
	var i4, j4, k4 = 10, 11, 12
	s4, s5, s6 := "string1", "string2", "string3"

	fmt.Println(i4, j4, k4, s5, s6, s4)

	var (
		j7  = 10
		j8  = 11
		s10 = "string1"
	)
	fmt.Println(j7, j8, s10)
}

/*
상수는 const 예약어를 통해 선언
*/
const i = 1
const s = "STRING"

func Constants() {
	fmt.Println(i, s)
}
