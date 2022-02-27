package p3

// Returns the largest prime factor of n
func largestPrimeFactor(n int64) int64 {
	var facs = primeFactors(n)
	// fmt.Println(facs)
	return facs[len(facs)-1]
}

// Returns a slice filled with all prime factors of n
func primeFactors(n int64) []int64 {
	var facs = make([]int64, 0)
	var i int64 = 2
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
func isPrime(n int64) bool {
	var i int64 = 2
	for ; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
