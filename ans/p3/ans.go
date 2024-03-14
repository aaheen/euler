package p3

import (
	"aaheen/euler/libeuler"
	"fmt"
)

// Returns solution for Problem 3
func Sol() {
	var pfacs []uint64 = libeuler.PrimeFacs(600851475143)
	fmt.Println(pfacs[len(pfacs)-1])
}

// Interactive solution session for Problem 3
func Ask() {
	fmt.Print("Find the largest prime factor of n. \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	var pfacs []uint64 = libeuler.PrimeFacs(n)
	fmt.Println("Answer:", pfacs[len(pfacs)-1])
}
