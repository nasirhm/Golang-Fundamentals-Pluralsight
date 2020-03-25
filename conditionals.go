package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main(){
	firstRank := "39"
	secondRank := "614"
	// If Statement : Boolean True or False Strict, We can nest ifs too.
	if firstRank < secondRank {
		fmt.Println("First course is doing better than second course")
	}else if firstRank > secondRank{
		fmt.Println("Oh Dear, The firstrank is greater than second")
	}else{
		fmt.Println("Something weird is going on")
	}
	// if <simple statement>; <Boolean Expression>
	if abc := 123; abc<322{
		// abc is constraint to the only line
		// as many else ifs only 1 else
		fmt.Println(abc)
	}

	// Switch
	// switch <simple statement>; <expression>
	// case <expression>: <code>
	// default: <code>
	// fallthrough, Executed the next ones too.
	switch "docker" {
	case "docker":
		fmt.Println("Docker")
		fallthrough
	case "linux":
		fmt.Println("Linux")
	case "Windows":
		fmt.Println("Windows")
	default:
		fmt.Println("We can't find any match")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got an even number -", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got an odd number -", tmpNum)
	}
	// Basic error handling with if
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}