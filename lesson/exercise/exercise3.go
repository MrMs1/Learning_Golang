package exercise

import "fmt"

type exVertex struct {
	X, Y int
}

func (v exVertex) Plus() int {
	return v.X + v.Y
}

func Exercise3_1() {
	v := exVertex{3, 4}
	fmt.Println(v.Plus())
}

func (v exVertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.X, v.Y)
}

func Exercise3_2() {
	v := exVertex{3, 4}
	fmt.Println(v)
}
