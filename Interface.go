package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s1 := square{10}
	c1 := circle{9}
	fmt.Println(s1.area())
	fmt.Println(c1.area())
}
