package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "This is channel demo which is very interesting"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	//Method 1 to pull from channel
	/* for i := 0; i < len(words); i++ {
		fmt.Print(<-ch, " ")
	} */

	//Method 2 to pull from channel
	for msg := range ch {
		fmt.Print(msg, " ")
	}
}