package main

import (
	"fmt"
	"os"
)

func main() {
	if a := 11; a>10 {
		fmt.Println("value is", a,"greater than 10")
	}

	// Read fallthrough
	switch name := "go is cool"; name {
	case "go is cool":
		fmt.Println(name)
	case "go is not cool":
		fmt.Println(name)
	default:
		fmt.Println("something weird")
	}

	_, err := os.Open("c:\\Users\\riteshagrahari")
	
	if err != nil {
		fmt.Println("Error occured", err)
	}

}