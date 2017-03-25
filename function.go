package main

import "fmt"

func main() {

	doSomething()
	add(1, 2)

}

func doSomething() {
	fmt.Println("Performing doSomething() function")
}

func add(a int, b int) {
	fmt.Println(a + b)
}
