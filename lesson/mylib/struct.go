package mylib

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func Structlesson() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)

	/*
		v := Vertex{x:1, y:2}
		fmt.Println(v)
		fmt.Println(v.x, v.y)
		v.x = 100
		fmt.Println(v.x, v.y)

		v2 := Vertex{x:1}
		fmt.Println(v2)

		v3 := Vertex{1, 2, "test"}
		fmt.Println(v3)
		fmt.Printf("v3: %T %v\n", v3, v3)

		v4 :=Vertex{}
		fmt.Println(v4)
		fmt.Printf("v4: %T %v\n", v4, v4)

		var v5 Vertex
		fmt.Println(v5)
		fmt.Printf("v5: %T %v\n", v5, v5)

		v6 := new(Vertex)
		fmt.Println(v6)
		fmt.Printf("v6: %T %v\n", v6, v6)

		v7 := &Vertex{}
		fmt.Printf("v7: %T %v\n", v7, v7)

		s := make([]int, 0)
		//s := []int{}
		fmt.Println(s)
	*/
}
