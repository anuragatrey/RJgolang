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

type shape interface {
	area() float64
}

func PrintArea(x shape) {
	fmt.Println("My area is :", x.area())
}

func main() {
	s1 := square{10}
	c1 := circle{9}
	PrintArea(s1)
	PrintArea(c1)
}
