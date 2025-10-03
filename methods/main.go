package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// value receiver - gets a copy of the struct
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

// pointer receiver - gets a reference to the struct
func (r *Rectangle) Double() {
	r.Width = r.Width * 2
	r.Height = r.Height * 2
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of rectangle:", Area(rect))
	rect.Double()
	fmt.Println("After doubling, area of rectangle:", Area(rect))
}
