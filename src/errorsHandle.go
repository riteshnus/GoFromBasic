package main

import (
	"fmt"
	"time"
	"os"
	"errors"
)

func main() {
	hoursOfDay := time.Now().Hour()
	greeting, err := getGreetings(hoursOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(greeting)
}

func getGreetings(hour int) (string, error) {
	var message string
	if hour < 9 {
		err := errors.New("Too early for greeting on weekend")
		return message, err
	}
	if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}
	return message, nil
}