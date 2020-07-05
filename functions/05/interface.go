package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (s square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	fmt.Printf("The area of the %T is %v\n", s, s.area())
}

func main() {
	s := square{
		length: 8,
	}
	c := circle{
		radius: 4,
	}
	info(s)
	info(c)
}
