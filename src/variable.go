package main

import (
	"fmt"
	"reflect"
)

var (
	name, course string
	module float64
	email = "abc@cde.com"
)
func main() {
	a := 10
	b := 3.2
	ptr := &b		//& refrence a pointer(pointer address) * de-refrence a pointer(value of pointer address)
	fmt.Println("value of A is", a, "and type is ", reflect.TypeOf(a))
	fmt.Println("value of name is", name, "and type is ", reflect.TypeOf(name))
	fmt.Println("value of module is", module, "and type is ", reflect.TypeOf(module))
	fmt.Println("value of email is", email, "and type is ", reflect.TypeOf(email))
	fmt.Println("pointer address of b is", ptr, "and value is ", *ptr)
}