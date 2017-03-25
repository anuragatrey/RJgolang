package main

import "fmt"

func main() {

	//If we put defer before any function it means that function will be executed
	// in the last before the main() function gets over
	defer doSomething()
	doSomethingElse()
}

func doSomething() {
	fmt.Println("doSomething()")
}

func doSomethingElse() {
	fmt.Println("doSomethingElse()")
}
