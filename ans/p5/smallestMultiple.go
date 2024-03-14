package p5

import "aaheen/euler/libeuler"

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
		case libeuler.IsPrime(i): // prime
			m *= uint64(i)
		case m%i != 0: // nonprime, not a factor of m already
			for j := uint64(2); j <= i>>1; j++ { // multiply by prime factors of i
				if i%j == 0 && libeuler.IsPrime(j) {
					m *= uint64(j)
				}
			}
		}
		// else: nonprime, already a factor of m
	}
	return m
}
