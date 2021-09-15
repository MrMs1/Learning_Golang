package mylib

import "fmt"

func one(x *int) {
	*x = 1
}

func Pointerlesson() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	fmt.Println(&*&n)

	fmt.Println(&n) // &でアドレス

	var p *int = &n // *で中身

	fmt.Println(p)

	fmt.Println(*p)

}

func NewMakelesson() {
	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	/*
		var p2 *int
		fmt.Println(p2)
		*p2++
		fmt.Println(p2)
	*/

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int) //値を入れる場合はmake
	fmt.Printf("%T\n", ch)

	var p3 *int = new(int) //ポインタを入れる場合はnew
	fmt.Printf("%T\n", p3)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
