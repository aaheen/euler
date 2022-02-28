package p4

import "fmt"

// Returns solution for Problem 4
func Sol() {
	n, _, _ := largestPalindromeProduct(3)
	fmt.Println(n)
}

// Interactive solution session for Problem 4
func Ask() {
	fmt.Print(`Find the largest palindrome made from the product of two n-digit numbers.
Input (0 < n < 10): `)
	var n int
	fmt.Scanln(&n)
	p, x, y := largestPalindromeProduct(n)
	fmt.Println("Answer:", p, "=", x, "*", y)
}
