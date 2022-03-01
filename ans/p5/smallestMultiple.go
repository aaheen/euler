package p5

// Returns smallest multiple of all numbers 1:n
func smallestMultiple(n int) (m uint64) {
<<<<<<< HEAD
	switch {
	case n < 1:
		panic("n must be at least 1")
	case n > 177: // I only know this number because of brute force testing
		panic("Answer Overflows uint64")
=======
	if n < 1 {
		panic("n must be at least 1")
>>>>>>> refs/remotes/origin/main
	}
	m = 1 // final product
	for i := 1; i <= n; i++ {
		switch {
		case prime(i): // prime
			m *= uint64(i)
		case m%uint64(i) != 0: // nonprime, not a factor of m already
			for j := 2; j <= i>>1; j++ { // multiply by prime factors of i
				if i%j == 0 && prime(j) {
					m *= uint64(j)
				}
			}
		}
		// else: nonprime, already a factor of m
	}
	return m
}

// Returns true if n is prime, false otherwise
func prime(n int) bool {
	for i := 2; i <= n>>1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
