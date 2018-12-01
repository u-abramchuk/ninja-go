package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area is", s.area())
}

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
	s := square{
		side: 5,
	}

	info(s)

	c := circle{
		radius: 5,
	}

	info(c)
}
