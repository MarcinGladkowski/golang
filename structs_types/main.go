package main

import (
	"fmt"
)

type ID int

type Person struct {
	Name string
	Age  int
}

type User struct {
	Person
	ID ID
}

func HelloPerson(u User) {
	fmt.Println(u)
}

func main() {

	var id ID = 1001

	p := Person{
		Name: "Gopher",
		Age:  10,
	}

	u := User{
		Person: p,
		ID:     id,
	}

	notInitialized := User{}

	HelloPerson(u)
	HelloPerson(notInitialized) // zeroes
}
