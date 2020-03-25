package main

import (
	"fmt"
	"time"
)

// Go only has one loop : for
// syntax : for <expression>
// Black expression = infinite loop

// self destruct
func main(){
	for timer := 10; timer >=0; timer--{
		if timer == 0{
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Millisecond)
	}

// for range loops
	// Slice of 3 Strings
	coursesInProg := []string{"Docker Deep Dive", "Docker Clustering", "Docker & Kubernetes"}
	coursesCompleted := []string{"Go Fundamentals", "Docker Deep Dive", "OS Fundamentals"}
	for _, i := range coursesInProg{
		fmt.Println(i)
		// Loop inside a loop
		for _, j := range coursesCompleted{
			if i==j {
				fmt.Println("Clash :",j)
			}
		}
	}

//	break & continue
	for timer := 10; timer >=0; timer-- {
		if timer % 2 == 0{
			// continue would reiterate the execution
			continue
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Millisecond)
	}

}