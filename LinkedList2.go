package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type student struct {
	name string
	age  int
	ssn  string
	next *student
}

func main() {

	students := new(student)
	students.next = nil

	if len(os.Args) < 2 {
		fmt.Println("There is no file")
		return
	}

	filename := os.Args[1]

	data, _ := ioutil.ReadFile(filename)

	for _, line := range strings.Split(string(data), "\n") {
		//Handle EOF
		if line == "" {
			continue
		}

		td := strings.Split(string(line), " ")
		ts := &student{
			name: td[0],
			ssn:  td[2],
			next: students.next,
		}
		//type cast the age in to Integer
		ts.age, _ = strconv.Atoi(td[1])

		students.next = ts

	}
	fmt.Println("List of our students")
	for s := students.next; s != nil; s = s.next {
		fmt.Println(s.name, s.age, s.ssn)
	}
}
