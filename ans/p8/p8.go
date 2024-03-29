package p8

import (
	euler "aaheen/euler/lib"
	"fmt"
	"os"
	"strconv"
)

// Returns solution for Problem 8
func Sol() {
	fmt.Println(productInStream(13, "ans/p8/p8.num"))
}

// Interactive solution session for Problem 8
func Ask() {
	fmt.Println("Find the n consecutive digits within x that have the greeatest product.")
	fmt.Print("Input n: ")
	var n uint64
	fmt.Scanln(&n)
	fmt.Print("Input path to x, stored as ASCII: ")
	var x string
	fmt.Scanln(&x)
	fmt.Println("Answer:", productInStream(n, x))
}

// Returns greatest product of n consecutive digits from stream x
func productInStream(n uint64, xpath string) uint64 {
	if n > 20 {
		fmt.Println("==== WARNING: n > 20 could overflow uint64 ====")
	}
	xfh, err := os.Open(xpath)
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s\n%s", xpath, err))
	}
	defer xfh.Close()

	bufBig := make([]byte, 32768) // 32kb should be enough
	bufSize, _ := xfh.Read(bufBig)
	var buf []byte = make([]byte, bufSize)
	copy(buf, bufBig)
	checkStreamValid(buf)

	var sByte []byte = gSumStreamSlice(n, buf)
	var s []uint64 = make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		s[i] = uint64(sByte[i])
	}

	return euler.SliceProd(s)
}

// Returns slice of n consecutive digits with greatest sum from stream s
func gSumStreamSlice(n uint64, s []byte) []byte {
	// immediately return s if n exceeds len(s)
	if n >= uint64(len(s)) {
		return s
	}

	var i, j, max uint64 = 0, 0, 0
	for j+n < uint64(len(s)-1) {
		if sum := uint64(euler.SliceSum(s[j : j+n])); !euler.HasZero(s[j:j+n]) && sum > max {
			i, max = j, sum
		}
	}

	return s[i : i+n]
}

// Panics if slice s has a non-numexical value
func checkStreamValid(s []byte) {
	for i := 0; i < len(s); i++ {
		_, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic("Non-numerical value in stream. Have you made sure there's no whitespace?")
		}
	}
	return
}
