package main

import "fmt"

func main() {

	//1. Way of looping
	for i := 0; i <= 10; i++ {
		fmt.Println("value of i:", i)
	}

	//2. way of looping - Infinite for loop also called for_while looping
	a := 0
	for a < 5 {
		fmt.Println("GOLANG")
		a++
		if a == 3 {
			break
		}
	}

	//3. way of looping
	myName := "Rishab"
	for _, Alpha := range myName {
		fmt.Print(string(Alpha) + " ")
	}
	fmt.Println()
}
