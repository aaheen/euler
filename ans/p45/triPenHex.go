package p45

// Returns slice of numbers that are at the intersection of
// the sets of all triagonal, pentagonal, and hexagonal nums.
// Generates each set's first n entries
func genIntx(n int) map[int]int {

	// generate sets
	t := make(map[int]int)
	p := make(map[int]int)
	h := make(map[int]int)
	for i := 0; i < n; i++ {
		t[tri(i)] = i
		p[pen(i)] = i
		h[hex(i)] = i
	}

	// find intersections
	k := make(map[int]int) // k for kernel i.e. intersection
	i := 0
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

func tri(n int) int {
	return n * (n + 1) / 2
}

func pen(n int) int {
	return n * (3*n - 1) / 2
}

func hex(n int) int {
	return n * (2*n - 1)
}
