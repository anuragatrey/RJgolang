package main

import "fmt"

func main() {

	//Variadic Functions -- Functions wich can take n Number of Arguments

	sum := add(5, 10, 15, 20, 25, 30, 50, 100, 45)
	fmt.Println(sum)
}

// ...int 3 dots before parameter type is used for defining variadic function
func add(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
