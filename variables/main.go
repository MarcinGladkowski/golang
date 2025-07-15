package main

func main() {
	var x int = 5
	var y float64 = 2.4
	z := 10 // equivalent to var z int = 10

	sumSameType := x + z
	sumConvertType := float64(x) + y

	println(sumSameType, sumConvertType)

	const pi float64 = 3.14
	const hello = "Hello" // no type declared

}
