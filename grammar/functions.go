package grammar

import (
	"fmt"
	"strings"
)

func Call() {
	sum(1, 2, 3, 4)

	prefix, suffix := divideStrings("Your_Name")
	fmt.Println(prefix, suffix)

	message := "Ronaldo"
	changeValue(&message) // 포인터 전달
	fmt.Println(message)

	// 익명 함수
	sum := func(x int, y int) int {
		return x + y
	}
	fmt.Println(sum(5, 10))

	// 일급 함수 , 함수 전달
	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(calc(add, 1, 2))
}

// 가변인자 ... values := []int{1, 2, 3, 4}
func sum(values ...int) int {
	result := 0
	for _, value := range values {
		result += value
	}

	return result
}

// 다중 변수 반환 : func 함수명(파라미터) (반환타입1, 반환타입2, ...) {}
func divideStrings(value string) (string, string) {
	values := strings.Split(value, "_")
	return values[0], values[1]
}

func changeValue(message *string) {
	*message = fmt.Sprintf("your name is %s", *message)
}

func calc(f func(int, int) int, a int, b int) int {
	return f(a, b)
}
