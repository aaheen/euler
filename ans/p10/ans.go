package p10

import "fmt"

// Returns solution for Problem 10
func Sol() {
	var sum uint64 = primeSum(2000000)
	fmt.Println("Sum of primes < 2M =", sum)
}

// Interactive solution session for Problem 10
func Ask() {
	fmt.Print("Find the sum of all primes < n.\nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Answer:", primeSum(n))
}
