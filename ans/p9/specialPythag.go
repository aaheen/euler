package p9

import (
	"fmt"
	"math"
)

// Returns pythagorean triplet that sums to N (if it exists) otherwise NULL
// This is the brute-force (dumb) implementation
func brutePythagTriN(n int) [3]int {
	for h := 1; h <= n; h++ {
	for j := h+1; j <= n; j++ {
		ksq := h*h + j*j
		if !isSquare(ksq) {
			continue
		} 
		k := int(math.Sqrt(float64(ksq)))
		if (sumTriple([3]int{h, j, k}) == n) {
			return [3]int{h, j, k}
			fmt.Println([3]int{h, j, k})
		}
	} }
	return [3]int{-1,-1,-1}
}

// Returns whether or not n is a square
func isSquare(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 1; i <= n/2; i++ {
		if i * i == n {
			return true
		}
	}

	return false
}

// Returns the product of triple t
func sumTriple(t [3]int) int {
	return t[0] + t[1] + t[2]
}

// Returns the product of triple t
func prodTriple(t [3]int) int {
	return t[0] * t[1] * t[2]
}

/*
func testIsSq() {
	fmt.Println("-1 is square:", isSquare(-1))
	fmt.Println("0 is square:", isSquare(0))
	fmt.Println("1 is square:", isSquare(1))
	fmt.Println("24 is square:", isSquare(24))
	fmt.Println("25 is square:", isSquare(25))
	fmt.Println("26 is square:", isSquare(26))
	fmt.Println("1023 is square:", isSquare(1023))
	fmt.Println("1024 is square:", isSquare(1024))
	fmt.Println("1025 is square:", isSquare(1025))
}
*/
