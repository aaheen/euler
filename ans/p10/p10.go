package p10

import (
	euler "aaheen/euler/lib"
	"fmt"
)

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

// Returns the sum of all primes less than or equal to n. Basic wrapper
// function that's just a summing for loop. See euler/prime.go and
// euler/basic.go for more information on how this is computed.
func primeSum(n uint64) uint64 {
	return euler.SliceSum(euler.Eratosthenes(n))
}
