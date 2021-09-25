package mylib

import "fmt"

const (
	c1 = iota
	c2
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func Iota() {
	StartLog()
	defer EndLog()
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}
