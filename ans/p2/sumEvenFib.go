package p2

// Returns the sum of all even fibonacci numbers below n
func sumEvenFib(n uint64) (sum uint64) {
	sum = 0
	var fibs = []uint64{1, 1, 2}
	for i := 2; fibs[i] <= n; fibs = append(fibs, fibs[i-1]+fibs[i-2]) {
		if fibs[i]%2 == 0 {
			sum += fibs[i]
		}
	}
	return sum
}
