package p5

import "fmt"

// Returns solution for Problem 5
func Sol() {
	fmt.Println(smallestMultiple(20))
}

// Interactive solution session for Problem 5
func Ask() {
	fmt.Print("Find the smallest multiple of all numbers 1:n. \nInput n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Answer:", smallestMultiple(n))
}
