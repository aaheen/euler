package p1

import "fmt"

// Prints solution for Problem 1
func Sol() {
	fmt.Println(sumThreeFive(1000))
}

// Interactive solution session for Problem 1
func Ask() {
	fmt.Print("Find the sum of all multiples of 3 or 5 below n. \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Answer:", sumThreeFive(n))
}

// Returns the sum of all multiples of 3 & 5 ≤ n
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
