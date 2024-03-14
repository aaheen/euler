package p7

import "fmt"

// Returns solution for Problem 7
func Sol() {
	fmt.Println(prime.numPrime(10001))
}

// Interactive solution session for Problem 7
func Ask() {
	fmt.Print("What is the nth prime number? \nInput n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Answer:", numPrime(n))
}
