package main

import "fmt"

// Maps are unordered list
// Maps are key:value pairs
// Maps are unsafe for concurrency
// make(map[<keyType>]<valueType>, size)
// maps are implementation of hashtable

func main() {
	// map[keyType]valueType, keyType should be comparable
	leagueTitles := make(map[string]int)
	leagueTitles["Sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle": 0,
	}

	fmt.Println(leagueTitles)
	fmt.Println(recentHead2Head)

	// Iterating maps
	testMap := map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5}
	for key, value := range testMap{
		fmt.Printf("Key is: %v value is: %v\n", key, value)
	}
	testMap["A"] = 100
	testMap["F"] = 173
	fmt.Println(testMap)
	delete(testMap, "F")
	fmt.Println(testMap)
}