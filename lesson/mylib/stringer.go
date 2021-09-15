package mylib

import (
	"fmt"
)

type Person2 struct {
	Name string
	Age  int
}

func (p2 Person2) String() string {
	return fmt.Sprintf("My name is %v", p2.Name)
}

func Stringerlesson() {
	mike := Person2{"Mike", 22}
	fmt.Println(mike)
}
