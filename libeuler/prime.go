package libeuler

import (
	"fmt"
	"math"
)

// Returns a slice of all prime numbers less than n. The initial size
// of the slice is determined using Pierre Dusart's equation for the
// upper bound of the prime counting function. (See PCUB() for more info.)
func Eratosthenes(n uint64) []uint64 {
	var pcub uint64 = PCUB(n)         // prime count upper bound
	var primes = make([]uint64, pcub) // slice of primes, i.e. return value
	// sieve = inverse truth table, each index contains the respective
	// !(primeness) of that number.
	var sieve = make([]bool, n+1)
	for i := uint64(2); i <= n; i++ {
		// If number is previously unmarked, it must be prime.
		if !sieve[i] {
			primes = append(primes, i)
		}
		// Mark all multiples of i as non-prime.
		for j := i; j <= n; j *= j {
			sieve[j] = true
		}
	}
	return primes
}

// Returns true if n is prime
func IsPrime(n uint64) bool {
	// test for divisibility by prime factors ≤ n / 2
	var testFacs = Eratosthenes(n >> 1)
	for i := uint64(0); i < uint64(len(testFacs)); i++ {
		if n%testFacs[i] == 0 {
			return false
		}
	}
	return true
}

// Returns the nth prime number. Computed using the Sieve of Eratosthenes.
func NthPrime(n uint64) uint64 {
	var nF64 float64 = float64(n)
	var upper float64 = nF64 * math.Log(nF64) * math.Log(nF64)
	for PCUB(uint64(upper)) < n {
		upper *= 1.1
		if upper >= math.MaxUint64 {
			panic(fmt.Sprint("Error: Size of sieve needed to compute nth prime exceeds MaxUint64. n =", n))
		}
	}
	return Eratosthenes(uint64(upper))[n-1]
}

// Prime Counting function Upper Bound estimate. Returns an upper bound
// estimate of the number of primes ≤ n. Computed using Pierre Dusart's
// equation estimating the upper bound of the prime counting function,
//
//	π(x) ≤ x / log(x) (1 + 1.2762 / log(x))
//
// This is chosen for the time being as it's just one line. More
// accurate but more complex functions will be implemented once I have
// the opportunity to actually understand how they work.
func PCUB(n uint64) uint64 {
	// Don't need to check log(n) ≤ maxUint64 as for large numbers log will always be ≤ n
	return n/uint64(math.Log(float64(n))) + uint64(1.0+(1.2762/math.Log(float64(n))))
}

// Returns a slice containing the prime factors of n
func PrimeFacs(n uint64) (facs []uint64) {
	var primes []uint64 = Eratosthenes(n >> 1)
	for i := uint64(0); i <= uint64(len(primes)); i++ {
		if n%primes[i] == 0 {
			facs = append(facs, primes[i])
			n /= primes[i]
		}
		if n == 1 {
			break
		}
	}
	return facs
}
