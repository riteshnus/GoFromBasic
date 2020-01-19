package main

import (
	"fmt"
)

type gopher struct {
	name string
	age int
	isAdult bool
}

func main() {
	var myName *string = new(string)
	*myName = "Ritesh"
	fmt.Println(*myName)

	name := "Nigel"
	course := "JavaScript: Basic"

	ptr := &name
	fmt.Println(ptr, *ptr)

	fmt.Println("\nHey",name,", You are running course", course)
	newCourse := changeCourse(course)
	fmt.Println("\nreturn course", newCourse)
	fmt.Println("\nYou are now watching course", course)

	phil := &gopher{name: "Phil", age : 37}
	fmt.Println("\nBefore",phil)
	validateAge(phil)
	fmt.Println("\nAfter",phil)
}

func changeCourse(course string) string {
	course = "Typescript: Deep dive"
	fmt.Println("\nCouse is changing to", course)
	return course
}

func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}