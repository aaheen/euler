package p8

import (
	"fmt"
	"os"
	"strconv"
)

// Returns greatest product of n consecutive digits from stream x
func productInSeries(n int, xpath string) uint64 {
	if n > 20 {
		fmt.Println("==== WARNING: n > 20 could overflow uint64 ====")
	}
	xfh, err := os.Open(xpath)
	if err != nil {
		panic(fmt.Sprintf("Unable to read file: %s\n%s", xpath, err))
	}
	defer xfh.Close()
	//fmt.Println("Opened", xpath)
	buf := make([]byte, 32768) // 32kb should be enough
	bufSize, _ := xfh.Read(buf)
	xx := make([]byte, bufSize)
	copy(xx, buf)
	//fmt.Println(copy(xx, buf), "bytes copied,",len(xx), "len")
	i := 0
	for j, max := 0, uint64(0); j+n < len(xx)-1; j++ {
		if p := prodSlice(xx[j : j+n]); p > max {
			i, max = j, p
		}
	}
	return prodSlice(xx[i : i+n])
}

// Returns slice of n consecutive digits with greatest sum from stream xx
func sumInSeriesNums(n int, xx []byte) []byte {
	if n >= len(xx) {
		return xx
	}
	i := 0
	for j, max := 0, 0; j+n < len(xx)-1; j++ {
		if sum := sumSlice(xx[j : j+n]); !hasZero(xx[j:j+n]) && sum > max {
			i, max = j, sum
		}
	}
	return xx[i : i+n]
}

// Returns sum of digits in slice xx
func sumSlice(xx []byte) int {
	sum := 0
	for i := 0; i < len(xx); i++ {
		//fmt.Printf("%s ", string(xx[i]))
		x, err := strconv.Atoi(string(xx[i]))
		if err != nil {
			panic("Non-numerical value in stream. Have you made sure there's no whitespace?")
		}
		sum += x
	}
	//fmt.Print("\n")
	return sum
}

// Returns product of digits in slice xx
func prodSlice(xx []byte) uint64 {
	var p uint64 = 1
	for i := 0; i < len(xx); i++ {
		x, err := strconv.Atoi(string(xx[i]))
		if err != nil {
			panic("Non-numerical value in stream. Have you made sure there's no whitespace?")
		}
		//fmt.Printf("%d ", x)
		p *= uint64(x)
	}
	return p
}

// Returns true if xx has a zero, false otherwise
func hasZero(xx []byte) bool {
	for i := 0; i < len(xx); i++ {
		x, err := strconv.Atoi(string(xx[i]))
		if err != nil {
			panic("Non-numerical value in stream. Have you made sure there's no whitespace?")
		}
		if x == 0 {
			return true
		}
	}
	return false
}
