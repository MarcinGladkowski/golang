package main

import (
	"fmt"
)

func modifyArray(arr [3]int) {
	arr[0] = 100
}

func modifySlice(arr []int) {
	arr[0] = 200
}

func main() {

	var a [5]int // [0, 0, 0, 0, 0] <- fixed size

	println(a[2])

	b := [3]int{1, 2, 3}

	fmt.Println(a)
	fmt.Println(b)

	c := [6]int{1, 2, 3, 4, 5, 6}
	d := c[1:4] // slice but still keep reference to the underlying array

	d[0] = 0 // d[0] is 2 before assign new value

	fmt.Println(c)

	// slice literals
	var e []int
	f := []int{1, 2, 3}
	fmt.Println(f)
	println(e == nil)

	// length
	s := [6]int{1, 2, 3, 4, 5, 6}
	t := s[2:4]

	fmt.Printf("len=%d cap=%d %v\n", len(t), cap(t), t)

	// another way to create a slice - allocate memory for zeroed slice references to an array
	x := make([]int, 5)
	y := make([]int, 0, 5)

	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
	fmt.Printf("len=%d cap=%d %v\n", len(y), cap(y), y)

	// append method
	n := []int{1, 2, 3}
	n = append(n, 4) // https://go.dev/blog/slices-intro
	fmt.Println(n)

	numbers := [3]int{10, 20, 30}

	for i, num := range numbers {
		fmt.Println(i, "=>", num)
	}

	// array is passed by value, not reference
	arrayToModify := [3]int{1, 2, 3}
	modifyArray(arrayToModify)
	fmt.Println(arrayToModify)

	sliceToModify := []int{1, 2, 3} // will be modified
	modifySlice(sliceToModify)
	fmt.Println(sliceToModify)
}
