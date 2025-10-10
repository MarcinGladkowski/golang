package main

import (
	"fmt"
	"sync"
)

func sum(s []int, result *int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	
	sum := 0
	for _, v := range s {
		sum += v
	}
	*result = sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	
	var wg sync.WaitGroup
	var x, y int
	
	// Add 2 goroutines to the wait group
	wg.Add(2)
	
	// Launch goroutines
	go sum(s[:len(s)/2], &x, &wg)
	go sum(s[len(s)/2:], &y, &wg)
	
	// Wait for both goroutines to complete
	wg.Wait()
	
	fmt.Println(x, y, x+y)
}