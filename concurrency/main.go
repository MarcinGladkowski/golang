package main

import (
	"fmt"
	"time"
)

func say(s string) {
	fmt.Println(s)
}

func main() {
	go say("world") // execute in separate goroutine
	say("hello")
	
	// We have added sleep to prevent program from exiting
	// before goroutine runs, there are better ways to
	// handle this using channels and wait groups
	time.Sleep(100 * time.Millisecond)
}