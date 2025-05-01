package grammar

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func IfCondition() {
	rand.New(rand.NewSource(time.Now().Unix())) // rand New 는 매번 다른 값을 주지 않으면 결과가 동일함
	i := rand.Intn(15)                          // 0부터 14까지 중 하나의 정수를 랜덤으로 생성

	fmt.Println(i)

	if i < 10 {
		fmt.Println("A")
	} else if i == 10 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// -------- if condition -----------
	str := "Hello World"
	if err := os.WriteFile("hello.txt", []byte(str), 0644); err != nil {
		fmt.Println(err)
	}

}

func Switch() {
	rand.New(rand.NewSource(time.Now().Unix()))
	i := rand.Intn(15)

	switch i {
	case 0, 1:
		fmt.Println("A")
	case 2, 3, 4:
		fmt.Println("B")
	case 5:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
