package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	// Your code here

	marcin := Person{Name: "Marcin"}
	other := Person{Name: "Marcin"}

	fmt.Println(marcin == other) // true

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr1 == arr2) // true

	// Slices, maps, and functions are not comparable: invalid operation
	tasks := []string{"Task1", "Task2", "Task3"}
	// tasks2 := []string{"Task1", "Task2", "Task3"}
	// fmt.Println(tasks == tasks2)

	fmt.Println(tasks != nil) // true

	// Pointers are compared by reference (address)
	pp1 := &Person{Name: "Bob"}
	pp2 := &Person{Name: "Bob"}
	pp3 := pp1

	fmt.Println("pp1 == pp2:", pp1 == pp2) // false - different instances
	fmt.Println("pp1 == pp3:", pp1 == pp3) // true - same instance
	fmt.Println("*pp1 == *pp2:", *pp1 == *pp2)
}
