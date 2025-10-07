package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printArea(s Shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	r := &Rectangle{length: 10, width: 5}
	c := &Circle{radius: 7}

	printArea(r)
	printArea(c)
}
