package mylib

import (
	"fmt"
)

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func (d Dog) Say() {
	fmt.Println(d.Name)
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func Interfacelesson() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"X"}
	var dog Dog = Dog{"dog"}
	DriveCar(mike)
	DriveCar(x)
	dog.Say()
	//DriveCar(dog)
}
