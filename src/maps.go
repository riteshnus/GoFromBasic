package main

import (
	"fmt"
)

func main() {
	// Maps are unordered list and call by refrence
	// Unsafe for concurrency
	// specify size can inprove performance: make(map[string]int, 100)
	winningTeams := make(map[string]int)
	winningTeams["Australia"] = 4
	winningTeams["India"] = 6

	recentHeadToHead := map[string]int{
		"England" : 5,
		"Westindies" : 3,
	}

	fmt.Printf("Leage teams:%v \nHead to Head:%v\n", winningTeams, recentHeadToHead)

	testMap := map[string]int{ "A" : 1, "B" : 2, "C" : 3}

	for key, value := range testMap {
		fmt.Printf("Key:%v - Value:%v\n", key, value)
	}

	testMap["A"] = 100
	fmt.Println(testMap)

	testMap["F"] = 200
	fmt.Println(testMap)

	delete(testMap, "B")
	fmt.Println(testMap)

	// composite way
	for key, value := range map[string]int {
		"X" : 1,
		"Y" : 2,
		"Z" : 3,
	} {
		fmt.Printf("Key:%v - Value:%v\n", key, value)
	}
}