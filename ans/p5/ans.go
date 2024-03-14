package p5

import "fmt"

// Returns solution for Problem 5
func Sol() {
	fmt.Println(leastFullComposite(20))
}

// Interactive solution session for Problem 5
func Ask() {
	fmt.Print("Find the smallest multiple of all numbers 1:n. \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Answer:", leastFullComposite(n))
}
