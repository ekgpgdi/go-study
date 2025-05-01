package grammar

import "fmt"

func For() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for {
		fmt.Println("HI")
		i++

		if i == 1 {
			break
		}
	}
}

// ForRange : for Range = foreach 형식
func ForRange() {
	strArray := []string{"A", "B", "C", "D", "E", "F", "G"}
	for i, str := range strArray {
		fmt.Println(i, str)
	}

	dictionary := map[string]string{
		"KeyA": "ValueA",
		"KeyB": "ValueB",
	}

	for key, value := range dictionary {
		fmt.Println(key, value)
	}
}

type Obj struct {
	Name string
	Age  int
}

func Range() {
	list := []Obj{
		{"A", 10},
		{"B", 20},
	}

	// _ : blank identifier(빈 식별자)
	for _, obj := range list {
		obj.Age = obj.Age * 2
		fmt.Println(obj.Name, obj.Age)
	}
}
