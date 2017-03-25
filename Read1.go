package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//As we are reading a flat file so need to define a structure
type student struct {
	name string
	age  int
	ssn  string
}

func main() {

	data, _ := ioutil.ReadFile("simple.txt")

	//Read Full File data
	fmt.Println("File Data.............")
	for _, read := range data {

		fmt.Print(string(read))
	}
	fmt.Println("End Of File Data......")

	for _, file_data := range strings.Split(string(data), "\n") {

		if file_data == "" {
			continue
		}

		td := strings.Split(string(file_data), " ")
		ts := &student{
			name: td[0],
			ssn:  td[2],
		}
		//type casting integers
		ts.age, _ = strconv.Atoi(td[1])
		fmt.Println()
		fmt.Println(td[0], td[1])
	}

}
