package main

import (
	"fmt"
	"strings"
)

/**
* The functions are higher order functions - can be passed as arguments and assigned to variables
 */
func hello(name string) string {
	if name != "" {
		name = "there"
	}

	return "Hello" + name
}

func makeMultiplier(multiplier int) func(int) int {
	return func(x int) int {
		return x * multiplier
	}
}

// example of returning multiple values, handy with returning errors
func parseName(fullName string) (string, string) {
	parts := strings.Split(fullName, ".")
	if len(parts) != 2 {
		return parts[0], ""
	}

	return parts[0], parts[1]
}

func main() {
	multiply := makeMultiplier(2)
	println(multiply(3))

	firstName, lastName := parseName("John Wick")

	fmt.Printf("%s, %s", firstName, lastName)
}
