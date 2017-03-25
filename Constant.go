package main

import "fmt"

func main() {

	const goldenRatio float64 = 1.618034 //Const is used for defining constant
	// Constant must always be intz with value else error
	//value of constant can not be changed

	fmt.Println(goldenRatio)

	//Constant Enumeration
	const (
		first = iota //iota is a function which is used to assign the constant
		//value to variables 0,1,2,3............
		second
		third
		fourth
	)
	fmt.Println(first, second, third, fourth)
	// Play with Iota
	const (
		A = iota * 5 //iota is a function which is used to assign the constant value to variables 0,1,2,3............
		B
		C
		D
	)
	fmt.Println(A, B, C, D)

	//goldenRatio = 6.46546 //It will return error as constant
	//value can not be changed
}
