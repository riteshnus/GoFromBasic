package main

import (
	"fmt"
)

func main() {

	myLearning := make([]string, 5, 10)
	fmt.Println(myLearning)
	fmt.Printf("length is %d \nCapacity is %d\n", len(myLearning), cap(myLearning))

	newWayForSlice := []string{"GO", "Docker", "Typescript"}
	fmt.Println(newWayForSlice)
	fmt.Printf("length is %d \nCapacity is %d\n", len(newWayForSlice), cap(newWayForSlice))


	// slice are passes as refrence to underline array
	mySlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)

	newSlice := make([]int, 1, 4)
	fmt.Printf("length is %d \nCapacity is %d\n", len(newSlice), cap(newSlice))

	for i := 1; i < 17; i++ {
		newSlice = append(newSlice, i)
		fmt.Printf("Capacity is %d\n", cap(newSlice))
	}

	sliceTest := []int{1, 2, 3, 4, 5, 6, 7}
	sliceAppend := []int{10, 20, 30}
	sliceTest = append(sliceTest, sliceAppend...)
	fmt.Println(sliceTest)
}