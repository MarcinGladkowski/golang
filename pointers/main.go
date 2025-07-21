package main

import "fmt"

type Object struct {
	Value int
}

func notModifyObject(object Object) {
	object.Value = 10
}

func modifyObject(object *Object) {
	object.Value = 10
}

func main() {

	//var i *int

	notModified := Object{Value: 1}

	notModifyObject(notModified)
	println("not modified object property:", notModified.Value)

	modifyObject(&notModified)
	println("modified object property:", notModified.Value)

	// i == nil // true

	x := 42
	i := &x

	fmt.Println(*i)

	*i = 82
	fmt.Println(x)

	// var y *string
	// fmt.Println(*y) -> panic: runtime error: invalid memory address or nil pointer dereference
}
