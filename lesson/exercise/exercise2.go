package exercise

import "fmt"

func Exercise2_1() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var min int
	for i := 0; i < len(l); i++ {
		if i == 0 {
			min = l[i]
			continue
		}
		if l[i] < min {
			min = l[i]
		}
	}
	fmt.Println("min value:", min)
}

func Exercise2_2() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var sum int
	for _, v := range m {
		//sum = sum + v
		sum += v
	}
	fmt.Println("sum:", sum)

}
