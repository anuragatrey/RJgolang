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
	next *student
}

func main() {

	data, _ := ioutil.ReadFile("simple.txt")

	//Read Full File data
	fmt.Println("File Data.............")
	for _, read := range data {

		fmt.Print(string(read))
	}
	fmt.Println("End Of File Data......")

	//If want to read some data from a file -- LinkedList
	students := new(student)
	students.next = nil

	for _, file_data := range strings.Split(string(data), "\n") {

		if file_data == "" {
			continue
		}

		td := strings.Split(string(file_data), " ")
		ts := &student{
			name: td[0],
			ssn:  td[2],
			next: students.next,
		}
		//type casting integers
		ts.age, _ = strconv.Atoi(td[1])
		students.next = ts

	}

	fmt.Println()
	fmt.Println("The data present in File simple.txt")
	fmt.Println("Column1 - Name", "Column2 - Age", "Column3 - SSN")

	for s := students.next; s != nil; s = s.next {
		fmt.Println(s.name, s.age, s.ssn)
	}

}
