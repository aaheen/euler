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
