package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 10; timer >= 0 ; timer -- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	learning := []string{"GO", "Docker", "Typescript"}
	completed := []string{"Javascript", "Typescript", "Cloud"}
	// planned := []string{"DCA", "Devops", "GO", "Data"}

	for _, i := range learning {
		// fmt.Println(i)
		for _, j := range completed {
			if i==j {
				fmt.Println("Hey Found language", j, "in both list")
			}
		} 
	}

	for timer := 10; timer >= 0 ; timer -- {
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

	// breakpoint to skip level of loop
	/* for _, i := range completed {
		completed 
		for _, j := range learning {
			for _, k := range planned {
				if k==j {
					fmt.Println("Hey Found language", j, "in planned and learning list")
					break completed
				}
			}
		} 
	} */
}