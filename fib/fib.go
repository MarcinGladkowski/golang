package fib

// Since this function doesn't start with a capital letter it isn't exported
func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// FibonacciSequence This function is exported since it starts with a capital letter
func FibonacciSequence(n int) []int {
	sequence := make([]int, n)

	for i := 0; i < n; i++ {
		sequence[i] = fibonacci(i)
	}

	return sequence
}
