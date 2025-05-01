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

	fmt.Println("==문자열 처리 (strings 패키지) ==")
	grammar.StringsPackage()
}
