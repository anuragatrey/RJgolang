package main

import "fmt"

func main() {
	//Slices are dynamic you can append/delete at any time & not in array
	// For Init slice use [] then define datatype of slice. eg, Strin, Int
	//Slice can have a single data Type & Not multiple.
	// Un-Initialized slice - {}
	//If want to initialize slice - []string{"Hello" , "Hi", "WhatsUp" }
	r := []string{}
	// making slice using inbuilt keyword slice "MAKE" - Make allocates memory
	// to each element of the slice. with make {} are not required
	// 3- denotes slice of 3 elements
	s := make([]string, 3)
	fmt.Println(r)
	fmt.Println(s)

	s[0] = "Hello"
	s[1] = "Hi"
	s[2] = "WhatsUp"

	fmt.Println(s[0], s[1], s[2])
}
