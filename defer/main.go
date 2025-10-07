package main

import (
	"fmt"
)

func riskyOperation() {
	defer func() {
	    if r := recover(); r != nil {
	        fmt.Println("Recovered from panic:", r)
	    }
	}()

	// This will cause a panic
	var arr []int
	fmt.Println(arr[1]) // Accessing out of bounds
}

func main() {
	riskyOperation()
	fmt.Println("Program continues after recovery")
}
