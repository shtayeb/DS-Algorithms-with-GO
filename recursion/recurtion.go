package recursion

func fibonacci(n int) int {
	// Base Case
	if n == 1 || n == 2 {
		return 1
	}

	if n <= 0 {
		return 0
	}

	// recurse
	return fibonacci(n-1) + fibonacci(n-2)
}

func factorial(n int) int {
	// base case
	if n == 1 || n == 0 {
		return 1
	}

	// recurse
	return n * factorial(n-1)
}
