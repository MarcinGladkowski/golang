package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type User struct {
	Person Person
	ID     string
}

func HelloPerson(u User) {
	fmt.Println(u)
}

func main() {

	p := Person{
		Name: "Gopher",
		Age:  10,
	}

	u := User{
		Person: p,
		ID:     "123-xyz",
	}

	HelloPerson(u)
}
