package mylib

import "fmt"

type Vertex2 struct {
	x, y int
}

// method
//　値レシーバー
func (v Vertex2) Area() int {
	return v.x * v.y
}

// ポインタレシーバー
func (v *Vertex2) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex23D struct {
	Vertex2
	z int
}

func (v Vertex23D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex23D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func Area(v Vertex2) int {
	return v.x * v.y
}

func New(x, y, z int) *Vertex23D {
	return &Vertex23D{Vertex2{x, y}, z}
}

func Methodslesson() {
	//v := Vertex2{3, 4}
	//fmt.Println(Area(v))
	v := New(3, 4, 5)
	v.Scale(10)
	fmt.Println(v.Area())
	fmt.Println(v.Area3D())

	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
}
