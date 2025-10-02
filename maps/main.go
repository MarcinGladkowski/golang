package main

import (
	"fmt"
)

type UserAge struct {
	age int
}

func main() {
	fmt.Println("Hello, World!")

	userScores := map[string]int{
		"Marcin": 36,
		"Tomek": 29,
		"Andrzej": 41,
	}

	fmt.Println("Marcin's score is:", userScores["Marcin"])

	userAges := make(map[string]UserAge)
	userAges["Marcin"] = UserAge{age: 36}
	userAges["Tomek"] = UserAge{age: 29}
	userAges["Andrzej"] = UserAge{age: 41}

	marcinAge := userAges["Marcin"]
	fmt.Println("Marcin's age is:", marcinAge)

	delete(userAges, "Tomek")

	fmt.Println(len(userAges))

	// accessing a non-existing key
	unknownAge, exists := userAges["Unknown"]
	fmt.Println("Unknown's age is:", unknownAge, "Exists:", exists) // returning 0, but second value exsists is false to indicate key not found
}