package main

import (
	"fmt"
)

type gopher struct {
	name string
	age int
}

type horse struct {
	name string
	weigth int
}

type jumper interface {
	jump() string
}

func (g gopher) jump() string {
	if g.age < 50 {
		return g.name + " can jump high"
	}
	return g.name + " can not jump"
}

func (h horse) jump() string {
if h.weigth > 2500 {
		return h.name + " is too heavy, can not jump high"
	}
	return h.name + " can jump easily"
}

func main() {
	jumperList := getList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.jump())
	}
}

func getList() []jumper{
	phil := gopher{name: "Phil", age : 30}
	noodles := gopher{name: "noodles", age : 90}
	barboro := horse{name: "barboro", weigth: 2000}

	list := []jumper{phil, noodles, barboro}
	return list
}