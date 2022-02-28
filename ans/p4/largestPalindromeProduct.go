package p4

import (
	"fmt"
	"strconv"
)

// Returns largest palindrome product made from two n-digit numbers,
// as well as what those two numbers actually are. p = x * y
func largestPalindromeProduct(n int) (p, x, y uint64) {
	switch {
	case n < 1:
		panic("Number of digits must be at least 1")
	case n > 9:
		panic("Cannot compute palindrome from numbers with more than 9 digits")
	case n > 6:
		fmt.Println("==== WARNING: Exceedingly slow for n > 6 ====")
	}
	// Declare bounds and default values
	var min, max uint64 = nines(n-1) + 1, nines(n)
	x, y = min, min
	p = x * y
	// Test all values
	for i := max; i >= min && i*max > p; i-- {
		for j := max; j >= min; j-- {
			if isPal(i*j) && (i*j) > p {
				p, x, y = i*j, i, j
			}
		}
	}
	return p, x, y
}

// Tests if n is a palidrome or not.
func isPal(n uint64) bool {
	var ns string = strconv.Itoa(int(n))
	for i := 0; i < len(ns)>>1; i++ {
		if ns[i] != ns[len(ns)-i-1] {
			return false
		}
	}
	return true
}

// Returns 10^d - 1
func nines(d int) uint64 {
	switch {
	case d < 0:
		panic("Cannot work with negative powers")
	case d == 0:
		return 0
	}
	var dd string = "9"
	for i := 1; i < d; i++ {
		dd = dd + "9"
	}
	i, err := strconv.Atoi(dd)
	if err != nil {
		panic(err)
	}
	return uint64(i)
}
