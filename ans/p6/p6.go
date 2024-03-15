package p6

import "fmt"

// Returns solution for Problem x
func Sol() {
	fmt.Println(diffSqSum(100))
}

// Interactive solution session for Problem x
func Ask() {
	fmt.Print(`Find the difference between
	The sum of the squares of the first n natural numbers, and
	The square of the sum.
Input n: `)
	var n int
	fmt.Scanln(&n)
	fmt.Println("Answer:", diffSqSum(n))
}

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
