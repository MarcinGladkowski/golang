package main

import (
	"fmt"
	
)

func main() {
	sequence := fib.FibonacciSequence(10)

	// This would result in an error
	// firstFibonacciNumber := initial.fibonacci(1)

	fmt.Println("Fibonacci sequence of first 10 numbers:")
	fmt.Println(sequence)
}
