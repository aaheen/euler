package main

// Returns the largest prime factor of n
func largestPrimeFactor(n uint64) uint64 {
	var facs = primeFactors(n)
	// fmt.Println(facs)
	return facs[len(facs)-1]
}

// Returns a slice filled with all prime factors of n
func primeFactors(n uint64) []uint64 {
	var facs = make([]uint64, 0)
	var i uint64 = 2
	for ; i < n/2; i++ {
		if n%i == 0 {
			if isPrime(i) {
				facs = append(facs, i)
				n /= i
			}
		}
	}
	if isPrime(n) {
		facs = append(facs, n)
	}
	return facs
}

// Returns true if n is prime, false otherwise
func isPrime(n uint64) bool {
	var i uint64 = 2
	for ; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
