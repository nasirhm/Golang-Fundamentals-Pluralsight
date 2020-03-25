package main

import (
	"fmt"
	"os"
	"reflect"
)

// 0 for int
// empty string for string
// global
var(
	name, course, module = "Nasir", "Go Fundamentals", 3.4
)

func main()  {
	instructorsName := "Nigel" // Only works within functions and declaring at the same time
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	// int(module) //for int type-conversion
	fmt.Println(instructorsName) // Variables must be used if only declared in function

	// Pointers & : references , * : dereferences, Go passes arguments by Value
	ptr := &module
	fmt.Println("Memory address of module variable is", &ptr,"and the value of the *module* is", *ptr)

	fmt.Println("The Course is", course)
	// No impact on the original value as go creates a copy of the variable when passing it to functions
	changeCourse(course)
	fmt.Println("The new course is", course)

	// Work around to impact the original variable
	changeCourses(&course)
	fmt.Println("the new value is", course)
	
	// Const, cant use `:=` convention
	const speedOfLight  = 186000
	fmt.Println("Speed of light is a constant and it is :",speedOfLight)

	// Accessing Environment variables in GO
	// Use the "os" Package
	// Environment Variables
	fmt.Println(os.Environ())

	// One per line
	for _, env := range os.Environ(){
		fmt.Println(env)
	}

	// assigning name the OS username
	name = os.Getenv("USERNAME")
	fmt.Println(name)
}

func changeCourse(c string) string {
	c = "First look: Go Lang Fundamentals"

	fmt.Println("Trying to change your course to", c)

	return c
}

func changeCourses(c *string) string {
	*c = "Go Lang Fundamentals"

	fmt.Println("Trying to change your course to", *c)

	return *c
}