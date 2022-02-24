package main

// Returns the sum of all even fibonacci numbers below a certain value
func SumEvenFib(limit int) int {
	var sum, n int = 0, 16
	// Is fib(n) enough? Check and gen. more if necessary
	for fib(n) < limit {
		n = n << 1
	}
	var fibs = fibList(n)
	// fibs now contains all fibonacci nums below limit, and most likely more
	for n = 0; fibs[n] <= limit; n++ {
		if fibs[n]%2 == 0 {
			sum += fibs[n]
		}
	}
	return sum
}

// Returns the nth fibonacci number
func fib(n int) int {
	return fibList(n)[n-1]
}

// Returns an array of the first n fibonacci numbers.
// Generated iteratively, NOT recursively
func fibList(n int) []int {
	var fibs = make([]int, n)
	for i := 0; i < n; i++ {
		switch i {
		case 0:
			fibs[0] = 1
		case 1:
			fibs[1] = 1
		default:
			fibs[i] = fibs[i-1] + fibs[i-2]
		}
	}
	return fibs
}

// Returns the nth fibonacci number, generated recursively.
// DANGEROUSLY slow for numbers larger than 30
func fibR(n int) int {
	switch {
	case n <= 0:
		return 0
	case n == 1:
		return 1
	default:
		return fibR(n-1) + fibR(n-2)
	}
}
