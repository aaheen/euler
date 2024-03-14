package p10

import "fmt"

// Returns solution for Problem 10
func Sol() {

	fmt.Println("Sum of primes < 2M")
}

// Interactive solution session for Problem 10
func Ask() {
	fmt.Print("Find the sum of all primes < n.\nInput n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Answer:", primeSum(n))
}
