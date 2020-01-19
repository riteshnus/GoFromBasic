package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello, I am running after 5 sec")
	}()

	// defer call when the function exit
	go func() {
		defer wg.Done()
		fmt.Println("I run immeditely")
	}()

	wg.Wait()
}