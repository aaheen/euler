// Basic utilities file. Sums slices, generates sequences, verifies
// Pythagorean triples, declares some custom constants, etc.

package euler

import (
	"math"
	"strconv"
)

// All number types (i.e. able to perform arithmetic)
type number interface {
	// byte ≡ uint8
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Returns a slice of the first n fibonacci numbers.
//
// Note: While n could be easily constrained to the range of 1 - 2^31
// (i.e. int) and save myself some type conversions when calling, I
// want to stay in line with the convention set by the rest of the
// project and use uint64 instead.
func Fibs(n uint64) []uint64 {
	var fibs = make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		switch i {
		case 0:
			fibs[0] = 1
		case 1:
			fibs[1] = 1
		default:
			fibs[i] = fibs[i-1] + fibs[i-2]
		}
	}
	return fibs
}

// Returns the nth fibonacci number, generated recursively.
// DANGEROUSLY slow for numbers larger than 30.
// Written just for fun. mmmmm toasty silicon
func FibRecursive(n int) int {
	switch {
	case n <= 0:
		return 0
	case n == 1:
		return 1
	default:
		return FibRecursive(n-1) + FibRecursive(n-2)
	}
}

// Returns true if n is a palidrome.
func IsPal(n uint64) bool {
	var ns string = strconv.Itoa(int(n))
	for i := 0; i < len(ns)>>1; i++ {
		if ns[i] != ns[len(ns)-i-1] {
			return false
		}
	}
	return true
}

// Returns true if n is even. If this was C I would have used a macro.
func IsEven(n uint64) bool {
	return n%2 == 0
}

// Returns true if slice s has a zero value, false otherwise
func HasZero[T number](s []T) bool {
	for i := uint64(0); i < uint64(len(s)); i++ {
		if s[i] == 0 {
			return true
		}
	}
	return false
}

// Returns the nth Fibonacci number. Uses the equation
//
//	Fₙ ≈ (φⁿ / √5)
//
// to approximate Fₙ, as Fₙ will always be the nearest whole number.
func NthFib(n uint64) uint64 {
	return uint64(math.Round(math.Pow(math.Phi, float64(n)/math.Sqrt(5))))
}

// Returns sum of all values in slice s.
func SliceSum[T number](s []T) (sum T) {
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}

// Returns product of all values in slice s.
func SliceProd[T number](s []T) (prod T) {
	for i := uint64(0); i < uint64(len(s)); i++ {
		prod += s[i]
	}
	return prod
}
