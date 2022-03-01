package p3

import "fmt"

// Returns solution for Problem 3
func Sol() {
	fmt.Println(largestPrimeFactor(600851475143))
}

// Interactive solution session for Problem 3
func Ask() {
	fmt.Print("Find the largest prime factor of n. \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Answer:", largestPrimeFactor(n))
}
