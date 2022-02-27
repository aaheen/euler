package p4

import "fmt"

// Returns solution for Problem 4
func Sol() {
	fmt.Println(largestPalindromeProduct(3))
}

// Interactive solution session for Problem 4
func Ask() {
	fmt.Print("Find the largest palindrome made from the product of two n-digit numbers.\nInput (n < 10): ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Answer:", largestPalindromeProduct(n))
}
