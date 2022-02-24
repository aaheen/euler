package main

import "fmt"

func main() {
	// Assuming user inputs valid num 0 : 18446744073709551615
	fmt.Print("Find the sum of all even fibonacci numbers below n. \nInput n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Sum:", SumEvenFib(n))
}
