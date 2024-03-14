package main

import "math"

// Returns a slice of all prime numbers less than n. The initial size
// of the slice is determined using Pierre Dusart's equation for the
// upper bound of the prime counting function. (See PCountUB()
// for more info.)
func Eratosthenes(n uint64) []uint64 {
	var pcub uint64 = PCountUB(n)     // prime count upper bound
	var primes = make([]uint64, pcub) // return value
	// sieve = inverse truth table, each index contains the respective
	// primeness of that number
	var sieve = make([]bool, n)
	var i, j uint64 = 2, 1
	for i <= n {
		if !sieve[i] {
			primes = append(primes, i)
		}
		for j = i; j <= n; {
			sieve[j] = true
			j *= j
		}
		i++
	}
	return nil
}

// Returns true if n is prime
func IsPrime(n uint64) bool {
	for i := uint64(2); i <= n>>1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Returns the nth prime number
func NthPrime(n uint64) uint64 {
	if n == 1 {
		return 2
	}
	var p uint64 = 1
	for i := 2; i <= n; {
		p += 2
		if IsPrime(p) {
			i++
		}
	}
	return p
}

// Returns the upper bound of the number of primes less than or equal
// to n. Computed using Pierre Dusart's equation estimating the upper
// bound of the prime counting function,
//
//	pi(x) ≤ x / log(x) (1 + 1.2762 / log(x))
//
// This is chosen for the time being as it's just one line. More
// accurate but more complex functions will be implemented once I have
// the opportunity to actually understand how they work.
func PCountUB(n uint64) uint64 {
	var x uint64 = uint64(n)
	return x/uint64(math.Log(float64(x))) + uint64(1.0+(1.2762/math.Log(float64(x))))
}
