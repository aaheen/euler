package p5

import (
	euler "aaheen/euler/lib"
	"fmt"
)

// Returns solution for Problem 5
func Sol() {
	fmt.Println(leastFullComposite(20))
}

// Interactive solution session for Problem 5
func Ask() {
	fmt.Print("Find the smallest multiple of all numbers 1:n. \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Answer:", leastFullComposite(n))
}

// Returns smallest number evenly divisible by all numbers 1:n
func leastFullComposite(n uint64) (m uint64) {
	switch {
	case n < 1:
		panic("n must be at least 1")
	case n > 177: // I only know this number because of brute force testing
		panic("Answer exceeds MaxUint64")
	}
	m = 1 // final product
	for i := uint64(1); i <= n; i++ {
		switch {
		case euler.IsPrime(i): // prime
			m *= uint64(i)
		case m%i != 0: // nonprime, not a factor of m already
			for j := uint64(2); j <= i>>1; j++ { // multiply by prime factors of i
				if i%j == 0 && euler.IsPrime(j) {
					m *= uint64(j)
				}
			}
		}
		// else: nonprime, already a factor of m
	}
	return m
}
