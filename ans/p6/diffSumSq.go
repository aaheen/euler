package p6

// Returns Sum(x=1:n{x})^2 - Sum(x=1:n{x^2})
func diffSqSum(n int) uint64 {
	return sqSum(n) - sumSq(n)
}

// Returns Sum(x=1:n{x})^2
func sqSum(n int) uint64 {
	var s uint64
	for i := 1; i <= n; i++ {
		s += uint64(i)
	}
	return s * s
}

// Returns Sum(x=1:n{x^2})
func sumSq(n int) uint64 {
	var s uint64
	for i := 1; i <= n; i++ {
		s += uint64(i * i)
	}
	return s
}
