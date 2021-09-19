package main

import (
	"fmt"
	"lesson/exercise"
	"lesson/mylib"
	"os/user"
	"time"
)

const Pi = 3.14

const (
	Name = "taro"
)

// initは初めに処理される特別な関数名
func init() {
	fmt.Println("------ init ------")
}

func test() {
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
	for _, param := range params {
		fmt.Println(param)
	}
}

func varLessen() {

	// ()でまとめて宣言可能
	var (
		i                  int     = 1
		f64                float64 = 1.2
		s                  string  = "test"
		t, f               bool    = true, false
		iDefault           int
		f64Default         float64
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

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

//main関数がなければ動作しない
func main() {
	test()
	fmt.Println("Hello world", "TEST") //カンマ区切りで文字列連結
	fmt.Println(time.Now())
	fmt.Println(user.Current())

	varLessen()

	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 200)
	fmt.Println(r3)

	// array
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println("a:", a)

	/* サイズが決まっているので下記のような記述はできない
	var b [2]int = [2]int{100, 200}
	b = append(b, 300)
	*/

	// slice
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("n:", n)
	fmt.Println("n:", n[2])
	fmt.Println("n:", n[2:4])
	fmt.Println("n:", n[:2])
	fmt.Println("n:", n[2:])
	fmt.Println("n:", n[:])
	n[2] = 100
	fmt.Println("n:", n)

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println("board:", board)
	n = append(n, 100)
	fmt.Println("n:", n)
	n = append(n, 100, 200, 300, 400)
	fmt.Println("n:", n)

	// slice make capacity
	m := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)
	m = append(m, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)
	m = append(m, 0, 0, 0, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(m), cap(m), m)

	l := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)

	k := make([]int, 0) // メモリ確保あり
	var o []int         // メモリ確保なし
	fmt.Printf("len=%d cap=%d value=%v\n", len(k), cap(k), k)
	fmt.Printf("len=%d cap=%d value=%v\n", len(o), cap(o), o)

	o = make([]int, 5)
	//o = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		o = append(o, i)
		fmt.Println(o)
	}
	fmt.Println(o)

	// map
	p := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(p)
	fmt.Println(p["apple"])
	p["banana"] = 300
	fmt.Println(p)
	p["new"] = 500
	fmt.Println(p)

	fmt.Println(p["nothing"])

	v, ok := p["apple"]
	fmt.Println(v, ok)

	v2, ok2 := p["nothing"]
	fmt.Println(v2, ok2)

	p2 := make(map[string]int)
	p2["pc"] = 5000
	fmt.Println(p2)

	/*
		var p3 map[string]int
		p3["pc"] = 5000
		fmt.Println(p3)
	*/

	// inner function
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(2)

	func(x int) {
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
	exercise.Exercise1()

	// Iflesson
	mylib.Iflesson()

	// Forlesson
	mylib.Forlesson()

	// Rangelesson
	mylib.Rangelesson()

	// Switchlesson
	mylib.Switchlesson()

	// Defer
	mylib.Deferlesson()

	// Logger
	//mylib.Loggerlesson()

	// ErrorHandling
	//mylib.ErrHandlinglesson()

	// PanicRecover
	mylib.PanicRecoverlesson()

	// exercise2
	exercise.Exercise2_1()
	exercise.Exercise2_2()

	// Pointerlesson
	mylib.Pointerlesson()

	// NewMakelesson
	mylib.NewMakelesson()

	// Structlesson
	mylib.Structlesson()

	// Methodslesson
	mylib.Methodslesson()

	// Interfacelesson
	mylib.Interfacelesson()

	// Typeassertionlesson
	mylib.TypeAssertionlesson()

	// Stringerlesson
	mylib.Stringerlesson()

	// CustomErrorlesson
	mylib.CustomErrorlesson()

	// exercise3
	exercise.Exercise3_1()
	exercise.Exercise3_2()

	// Goroutine
	mylib.Goroutine()

	// Channellesson
	mylib.Channellesson()

	// BufferedChannelslesson
	mylib.BufferedChannelslesson()

	//channellesson2
	mylib.Channellesson2()

	//producer&consumer
	mylib.ProducerConsumerlesson()

	// FanoutFaninlesson
	mylib.FanoutFaninlesson()
}
