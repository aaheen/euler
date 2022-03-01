package p7

// Returns the nth prime number
func numPrime(n int) uint64 {
	if n == 1 {
		return 2
	}
	var p uint64 = 1
	for i := 2; i <= n; {
		p += 2
		if prime(p) {
			i++
		}
	}
	return p
}

// Returns true if n is prime, false otherwise
func prime(n uint64) bool {
	for i := uint64(2); i <= n>>1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
