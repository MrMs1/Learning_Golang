package mylib

import "fmt"

func Exercise1() {
	f := 1.11
	fmt.Println("Q1.f convert to int:", int(f))
	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("Q3.%T %v\n", m, m)
}
