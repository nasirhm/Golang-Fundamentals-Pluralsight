package main

import "fmt"

// Structs Custom Data type
// Collections of fields

func main() {
	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}
	// var DockerDeepDive courseMeta
	// new keyword gives a pointer
	// DockerDeepDive := new(courseMeta)
	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}
	DockerDeepDive.Rating = 6
	fmt.Println("\nDocker Deep Dive's author is:", DockerDeepDive.Author)
	fmt.Println("\nDocker Deep Dive's rating is:", DockerDeepDive.Rating)

}