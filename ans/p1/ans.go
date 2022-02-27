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
