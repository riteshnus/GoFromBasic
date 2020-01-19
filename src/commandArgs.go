package main

import (
	"fmt"
	"os"
)

// command to run with args go run commandArg.go "I am learning"
// command to run without args goimports run commandArg.go "I am learning"
func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("I am Gopher")
	}
}