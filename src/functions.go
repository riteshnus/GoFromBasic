package main

import (
	"fmt"
)

func main() {
	lowestNumber := findLowest(13, 12, 2, 7, 8 , 3, 19, 15, 18)
	fmt.Println(lowestNumber)
}

func findLowest(number ...int) int {
	low := number[0]
	for _, i := range number {
		if i < low {
			low = i
		}
	}
	return low
}