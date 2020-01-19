package main

import (
	"fmt"
	"sync"
	"math"
)

//to run the code and get the execution time: time go run program1.go
//to run the code with single procs: time GOMAXPROCS=1 go run program1.go

/* concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once. */


/* 
goroutines are not OS Thread
	--> goroutines is LightWeight(work on top of OS Thread)
	--> Go manages goroutines(OS manages the Thread)
	--> go manages goroutines by go runtime(process thread are manage by os)
	--> Less Switching(when one goroutines sleep another task run on the same goroutines)
  --> faster startup time
	--> Safe communication through channel, it take care of responsibility of synchronization of Actor
*/

/* 
Go concurrency model is 
	--> Actor model
	--> Communicating Sequential processes (CSP)
*/

/* 
	--> goroutines are virtual thread
	--> 
 */
func main()  {
	names := []string{"Ritesh", "Akanksha", "Ashish"}
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names {
		go printName(name, &wg)
	}
	wg.Wait()
}

func printName(name string, wg *sync.WaitGroup) {
	result := 0.0
	for i:= 0; i<100000000; i++ {
		result += math.Pi * math.Sin(float64(len(name)))
	}
	fmt.Println("Name: ", name)
	wg.Done()
}