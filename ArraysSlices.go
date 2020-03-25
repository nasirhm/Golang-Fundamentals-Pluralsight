package main

import "fmt"

func main() {
	// Array has a fixed length
	// Slices, Variable length. Slices are built on top of arrays, They are passed by references to the array
	// make(<type>, <len>, <caps>)
	myCourses := make([]string, 5, 10)
	completedCourses := []string{"Docker", "Puppet", "Python"}
	fmt.Printf("Length is: %d. \nCapacity is: %d", len(myCourses), cap(myCourses))
	fmt.Printf("\nLength is: %d. \nCapacity is: %d\n", len(completedCourses), cap(completedCourses))
	fmt.Println(completedCourses[2])
	completedCourses[1] = "OpenShift"
	fmt.Println(completedCourses)

	// [:5] = [0:5]
	// [2:] = [2: len of array - 1]
	// 0 and 1
	sliceOfSlice := completedCourses[0:2]
	fmt.Println(sliceOfSlice)

	// Appending
	mySlice := make([]int, 1, 4)
	fmt.Printf("Length is: %d Capacity is: %d", len(mySlice), cap(mySlice))
	for i:=1; i<17 ;i++  {
		mySlice = append(mySlice, i)
		fmt.Printf("\nCapacity is: %d \n", cap(mySlice))
	}
	// Appending slice with slice
	newSlice := []int{10,20,30}
	// Gives an error mySlice = append(mySlice, newSlice)
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}