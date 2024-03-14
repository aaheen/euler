package p10

import "aaheen/euler/libeuler"

// Returns the sum of all primes less than or equal to n. Basic wrapper
// function that's just a summing for loop. See libeuler/prime.go and
// libeuler/basic.go for more information on how this is computed.

func primeSum(n uint64) uint64 {
	return libeuler.SliceSum(libeuler.Eratosthenes(n))
}
