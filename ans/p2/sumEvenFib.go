package p2

// Returns the sum of all even fibonacci numbers below a certain value
func sumEvenFib(limit uint64) uint64 {
	var sum uint64 = 0
	var n int = 16
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

// Returns the sum of all even fibonacci numbers below a certain value
// without helper functions
func sumEvenFib2(limit uint64) (sum uint64) {
	sum = 0
	var fibs = []uint64{1, 1, 2}
	for n := 2; fibs[n] <= limit; fibs = append(fibs, fibs[n-1]+fibs[n-2]) {
		if fibs[n]%2 == 0 {
			sum += fibs[n]
		}
	}
	return sum
}

// Returns the nth fibonacci number
func fib(n int) uint64 {
	return fibList(n)[n-1]
}

// Returns a slice of the first n fibonacci numbers.
// Generated iteratively, NOT recursively
func fibList(n int) []uint64 {
	var fibs = make([]uint64, n)
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
