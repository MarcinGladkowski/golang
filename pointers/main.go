package main

import "fmt"

func main() {
	var i *int

	// i == nil // true

	x := 42
	i := &x

	fmt.Println(*i)
}
