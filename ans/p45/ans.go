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
