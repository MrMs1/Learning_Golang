package mylib

import (
	"fmt"
	"sort"
)

func Sort() {
	StartLog()
	defer EndLog()

	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}

	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name }) //アルファベット順
	fmt.Println(i, s, p)
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age }) //年齢順
	fmt.Println(i, s, p)

}
