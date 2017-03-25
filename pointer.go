package main

import "fmt"

func main() {

	x := 1
	y := &x
	*y = x
	fmt.Println("The value of x is:", x)
	fmt.Println("The address of x is:", &x)
	fmt.Println("The value of y is:", *y)
	fmt.Println("The address of x is:", &y)
	*y = 100 //de-referencing

	fmt.Println("The value of x is:", x)
	fmt.Println("The value of y is:", y)
}
