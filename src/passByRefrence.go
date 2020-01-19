package main

import (
	"fmt"
	"os"
)

func main() {
	// name := "Nigel"
	name := os.Getenv("USERNAME")
	course := "JavaScript: Basic"

	fmt.Println("\nHey",name,", You are running course", course)
	changeCourse(&course)
	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {
	*course = "Typescript: Deep dive"
	fmt.Println("\nCouse is changing to", *course)
	return *course
}