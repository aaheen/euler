package p2

import "fmt"

// Returns solution for Problem 2
func Sol() {
	fmt.Println(sumEvenFib(4000000))
}

// Interactive solution session for Problem 2
func Ask() {
	fmt.Print("Find the sum of all even fibonacci numbers below n. \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Sum:", sumEvenFib(n))
}
