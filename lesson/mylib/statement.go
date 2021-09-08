package mylib

import (
	"fmt"
	"os"
	"time"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func Iflesson() {
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	/*
		num := 6
		if num % 2 == 0 {
			fmt.Println("by 2")
		} else if num % 3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	x, y := 11, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}

}

func Forlesson() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)

		sum := 1
		for sum < 10 {
			sum += sum
			fmt.Println(sum)
		}
		fmt.Println(sum)

		/* 無限ループ
		for {
			fmt.Println("hello")
		}
		*/
	}

}
func Rangelesson() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

}

func getOsName() string {
	return "mac"
}

func Switchlesson() {
	//os := getOsName()
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func Deferlesson() {
	foo()

	defer fmt.Println("world")

	fmt.Println("hello")

	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

func Loggerlesson() {

}
