package main

func SumThreeFive(n int) int {
	var sum int = 0
	for i := 0; i < n; i++ {
		switch {
		case i%3 == 0:
			sum += i
		case i%5 == 0:
			sum += i
		}
	}
	return sum
}
