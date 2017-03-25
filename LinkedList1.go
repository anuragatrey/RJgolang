package main

import "fmt"

// teacher(debrah) -> Student(John) -> student(Kyle)
type student struct {
	age    int
	weight int
	name   string
	next   *student
}

type teacher struct {
	age    int
	weight int
	name   string
	next   *student
}

func main() {
	kyle := student{7, 25, "Kyle", nil}
	john := student{8, 27, "John", &kyle}

	debrah := teacher{25, 48, "Debrah", &john}

	fmt.Println("The teacher name is:", debrah.name)
	fmt.Println("The student directly behind debrah is:", debrah.next.name)
	fmt.Println("The student directly behind John is:", debrah.next.next.name)
	fmt.Println("The student directly behind John is:", kyle.next)
}
