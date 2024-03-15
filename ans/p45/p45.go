package p45

import "fmt"

// Returns solution for Problem 45
func Sol() {
	k := genIntx(999999)
	for _, val := range k {
		fmt.Println(val)
	}
	fmt.Println()
}

// Interactive solution session for Problem 45
/* func Ask() {
	fmt.Print("Goal \nInput n: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Answer:", func(n))
} */

// Returns slice of numbers that are at the intersection of
// the sets of all triagonal, pentagonal, and hexagonal nums.
// Generates each set's first n entries
func genIntx(n uint64) map[byte]uint64 {

	// generate sets
	t := make(map[uint64]byte)
	p := make(map[uint64]byte)
	h := make(map[uint64]byte)
	for i := 0; uint64(i) < n; i++ {
		t[tri(i)] = byte(0)
		p[pen(i)] = byte(0)
		h[hex(i)] = byte(0)
	}

	// find intersections
	k := make(map[byte]uint64) // k for kernel i.e. intersection
	i := byte(0)
	for val := range h {
		_, isTri := t[val]
		_, isPen := p[val]
		if isTri && isPen {
			k[i] = val
			i++
		}
	}

	return k
}

// Returns the nth triagonal number
func tri(n int) uint64 {
	un := uint64(n)
	return un * (un + 1) / 2
}

// Returns the nth pentagonal number
func pen(n int) uint64 {
	un := uint64(n)
	return un * (3*un - 1) / 2
}

// Returns the nth hexagonal number
func hex(n int) uint64 {
	un := uint64(n)
	return un * (2*un - 1)
}
