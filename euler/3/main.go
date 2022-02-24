package main

import "fmt"

func main() {
	fmt.Print("Find the largest prime factor of n. \nInput n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Println("Largest prime factor:", largestPrimeFactor(n))
}
