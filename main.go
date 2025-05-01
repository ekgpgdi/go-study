package main

import (
	"fmt"
	"study/grammar"
)

func main() {
	fmt.Println("hello world")

	fmt.Println("==변수 학습==")
	grammar.Varivales()

	fmt.Println("==상수 학습==")
	grammar.Constants()

	fmt.Println("==문자열 포맷팅 학습==")
	grammar.StringFormat()

	fmt.Println("==문자열 처리 학습 (strings 패키지) ==")
	grammar.StringsPackage()

	fmt.Println("==문자열 -> 정수형 변환 학습 ==")
	grammar.StringToInt()

	fmt.Println("==if문 학습==")
	grammar.IfCondition()

	fmt.Println("==스위치 학습==")
	grammar.Switch()
}
