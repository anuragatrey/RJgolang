package main

import "fmt"

//Globle/Package level variables need to be declared as int in long form
//"Short way will not work"
var iron = 25

func main() {

	gold := 100

	fmt.Println("Scope of Main Function")
	fmt.Println(iron) // works -  Global variables are access anywhere
	fmt.Println(gold) // works too - Local varaibles defined can be access
	//in same function & not outside
	anotherFunction()
}

func anotherFunction() {

	fmt.Println("Scope of Another Function")
	//gold :=78
	fmt.Println(iron) //Works
	//fmt.Println(gold) // It wont work (Local variable accessed
	//outside main function)
}
