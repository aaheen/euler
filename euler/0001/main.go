package main

import "fmt"

func main() {
	fmt.Print("Find the sum of all multiples of 3 or 5 below n. \nInput n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Sum:", SumThreeFive(n))
}
