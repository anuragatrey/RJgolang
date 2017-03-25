package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//IO/IOUTIL is used for reading a file
	// Read the whole file at once
	b, err := ioutil.ReadFile("name.txt")

	if err != nil {
		panic(err)
		//Panic will not terminate the pgm abnormally if there is any error in pgm
		//It will log the error and come out of pgm
		//Tracks runTime Errors
	}
	for _, print_name := range b {
		fmt.Print(string(print_name))
	}
}
