package p7

import (
	euler "aaheen/euler/lib"
	"fmt"
)

// Returns solution for Problem 7
func Sol() {
	fmt.Println(euler.NthPrime(10001))
}

// Interactive solution session for Problem 7
func Ask() {
	fmt.Print("What is the nth prime number? \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Answer:", euler.NthPrime(n))
}
