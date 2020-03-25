package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "nigel poulton"
	bestFinish := bestLeagueFinishes(13, 10, 21, 24, 23, 21, 7, 3, 2)
	fmt.Println(bestFinish)
	fmt.Println(converter(module, author))
}

func converter(m string, author string) (s1, s2 string) {
	m = strings.Title(m)
	author = strings.ToUpper(author)
	return m, author
}

// Variadic Function : Function with unknown number of Parameters
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes{
		if i<best{
			best = i
		}
	}
	return best
}