package main

import "fmt"

func main() {
	var myAge = 25
	//another way of declaring the variable -No need of defining variable type,
	//It take type as the assigned value.
	// := is used for variable declaration
	myVariable := 37

	wAge, kAge := 28, 1

	fmt.Println(myAge)
	fmt.Println(wAge)
	fmt.Println(kAge)
	fmt.Println(myVariable)

	fmt.Println("--------Modified-------")

	// = sign is used to assign value to a variable
	myAge = 31
	myVariable = 50

	fmt.Println(myAge)
	fmt.Println(myVariable)

}
