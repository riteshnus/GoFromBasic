package model

type gopher struct {
	name string
	age int
}

type horse struct {
	name string
	weigth int
}

type jumper interface {
	Jump() string
}

func (g gopher) Jump() string {
	if g.age < 50 {
		return g.name + " can jump high"
	}
	return g.name + " can not jump"
}

func (h horse) Jump() string {
if h.weigth > 2500 {
		return h.name + " is too heavy, can not jump high"
	}
	return h.name + " can jump easily"
}

func GetList() []jumper{
	phil := gopher{name: "Phil", age : 30}
	noodles := gopher{name: "noodles", age : 90}
	barboro := horse{name: "barboro", weigth: 2000}

	list := []jumper{phil, noodles, barboro}
	return list
}