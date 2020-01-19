package main

import (
	"fmt"
)

type gopher struct {
	name string
	age int
}

func (g gopher) jump() string {
	if g.age < 40 {
		return g.name + " can jump HIGH."
	} 
		return g.name + " can still jump high."
}

func main() {

	type person struct {
		name string
		gender string
		age int
	}

	var Person1 person
	fmt.Println(Person1)

	// refer the pointer
	Person2 := new(person)
	fmt.Println(Person2)

	Person3 := person{
		name: "Ritesh",
		gender: "Male",
		age: 28,
	}
	fmt.Println(Person3)

	fmt.Println("\nName of the Author is:", Person3.name)

	gopher1 := gopher{name: "gopher1", age : 34}
	fmt.Println(gopher1.jump())
	gopher2 := gopher{name: "gopher2", age : 50}
	fmt.Println(gopher2.jump())
}