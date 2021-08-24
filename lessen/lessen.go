package main

import (
	"fmt"
	"os/user"
	"time"
)

const Pi = 3.14

const (
	Name = "taro"
)

// initは初めに処理される特別な関数名
func init()  {
	fmt.Println("------ init ------")
}

func test()  {
	fmt.Println("test")
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func hoge(params ...int) {
	fmt.Println("length:", len(params), "value:", params)
	for _, param := range params{
		fmt.Println(param)
	}
}

func varLessen(){

	// ()でまとめて宣言可能
	var (
		i int = 1
	 	f64 float64 = 1.2
	 	s string = "test"
	 	t,f        bool = true, false
		iDefault   int
		f64Default        float64
		sDefault           string
		tDefault, fDefault bool
	)

	fmt.Println(i, f64, s, t, f)
	fmt.Println(iDefault, f64Default, sDefault, tDefault, fDefault)

	// :=で型名の省略が可能
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
	fmt.Println(Pi, Name)
}

func add(x, y int) (int, int){
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}
//main関数がなければ動作しない
func main() {
	test()
	fmt.Println("Hello world","TEST") //カンマ区切りで文字列連結
	fmt.Println(time.Now())
	fmt.Println(user.Current())

	varLessen()

	r1, r2 := add(10,20)
	fmt.Println(r1, r2)

	r3 := cal(100, 200)
	fmt.Println(r3)

	// array


	// inner function
	f := func(x int){
		fmt.Println("inner func", x)
	}
	f(2)

	func (x int){
		fmt.Println("inner func", x)
	}(2)

	// closure
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	c2 := circleArea(3)
	fmt.Println(c2(2))

	// Variadic parameter
	hoge(10, 20)
	hoge(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)
	hoge(s...)

	// exercise
	exercise1()
}


