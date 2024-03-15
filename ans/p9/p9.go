package p9

import (
	"fmt"
	"math"
)

// Returns solution for Problem 9
func Sol() {
	tri := brutePythagTriN(1000)
	sum := sumTriplet(tri)
	prod := prodTriplet(tri)
	fmt.Println(tri, "Sum:", sum, "Product:", prod)
}

type triplet [3]uint64

var _NON_TRIPLET = triplet{0, 0, 0}

// Interactive solution session for Problem 9
func Ask() {
	fmt.Print("Find the pythagorean triple that sums to n (if it exists).\nInput n: ")

	var n uint64
	fmt.Scanln(&n)
	tri := brutePythagTriN(n)
	sum := sumTriplet(tri)
	prod := prodTriplet(tri)

	if tri == _NON_TRIPLET {
		fmt.Println("No triple with sum", n)
	} else {
		fmt.Println(tri, "Sum:", sum, "Product:", prod)
	}
}

// Returns pythagorean triplet that sums to N (if it exists) otherwise NULL
// This is the brute-force (dumb) implementation
func brutePythagTriN(n uint64) triplet {
	for h := uint64(1); h <= n; h++ {
		for j := h + 1; j <= n; j++ {
			var ksq = h*h + j*j
			if !isSquare(ksq) {
				continue
			}
			var k = uint64(math.Sqrt(float64(ksq)))
			var t = triplet{h, j, k}
			if sumTriplet(t) == n {
				return t
			}
		}
	}
	return triplet{0, 0, 0}
}

// Returns whether or not n is a square
func isSquare(n uint64) bool {
	if n <= 1 {
		return false
	}

	for i := uint64(1); i <= n/2; i++ {
		if i*i == n {
			return true
		}
	}

	return false
}

// Returns the product of triple t
func sumTriplet(t triplet) uint64 {
	return t[0] + t[1] + t[2]
}

// Returns the product of triple t
func prodTriplet(t triplet) uint64 {
	return t[0] * t[1] * t[2]
}
