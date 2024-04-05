package main

// n = 0 -> 0 // initial grandparent
// n = 1 -> 1 // initial parent
// n = 2 -> 1
// n = 3 -> 2
// n = 4 -> 3
// n = 5 -> 5
// n = 6 -> 8
// ....
func Fibonacci(n int) int {	
	if n <= 1 {
		// n = 0 -> 0
		// n = 1 -> 1
		return n
	}

	grandparent := 0
	parent := 1

	current := 0
	for i := 2; i <= n; i++ {
		current = parent + grandparent
		grandparent = parent
		parent = current
	}

	return current
}