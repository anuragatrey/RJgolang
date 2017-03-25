package main

import "fmt"

func main() {

	Greeting("Hello:", "James", "Katrina")
}

// ...int 3 dots before parameter type is used for defining variadic function
func Greeting(Prefix string, who ...string) {
	fmt.Println(Prefix)
	fmt.Println(who)
	for _, value := range who {
		fmt.Println(value)
	}
}
