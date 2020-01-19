package main

import (
	"fmt"
	"interface/model"
)

func main() {
	jumperList := model.GetList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}
