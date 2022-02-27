package p1

func sumThreeFive(n uint64) uint64 {
	var sum uint64 = 0
	for i := uint64(0); i < n; i++ {
		switch {
		case i%3 == 0:
			sum += i
		case i%5 == 0:
			sum += i
		}
	}
	return sum
}
