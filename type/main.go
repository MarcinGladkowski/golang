package main

type ID int

type Status string

const (
	StatusPending Status = "pending"
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type Response struct {
	Status  Status
	Message string
}

func main() {
	var defaultUserId ID = 10

	println(defaultUserId)

	errorResponse := Response{StatusError, "error"}

	println(errorResponse.Message)

	// compiles fine
	var s Status
	s = "some string" // Go Types are only aliases, we can assign any string for this variable with Status type

	println(s)
}
