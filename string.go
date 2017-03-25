package main

import "fmt"

func main() {

	//Concatenation
	string1 := "Hello Mr.Jain "
	string2 := "How are u ?"

	fmt.Println(string1 + string2)

	//Slicing Of String

	aSliceOfText := string1[:5]
	fmt.Println(aSliceOfText)

	name := string1[6:]
	fmt.Println(name)

}
