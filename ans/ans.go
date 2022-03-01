package ans

import (
	"fmt"
	"strconv"

	"github.com/pkg/browser"
)

// Prints out solution for problem p
func Sol(p int) {
	x := 1
	defer recAns(p, &x)
	var fn func() = solList[p]
	fn()
	x = 2
}

// Runs interactive solution for problem p
func Ask(p int) {
	x := 1
	defer recAns(p, &x)
	var fn func() = askList[p]
	fn()
	x = 2
}

// Opens problem p in my repo on Github
func Page(p int) {
	var gURL string = "https://github.com/m9ple/euler/tree/main/ans/p" + strconv.Itoa(p)
	fmt.Println("Heading to", gURL)
	browser.OpenURL(gURL)
}

// Opens my repo's main page on Github
func Repo() {
	var gURL string = "https://github.com/m9ple/euler"
	fmt.Println("Heading to", gURL)
	browser.OpenURL(gURL)
}

// Recovers if Sol or Ask fail
func recAns(p int, x *int) {
	if err := recover(); err != nil {
		fmt.Println("recAns", *x)
		switch *x {
		case 1:
			panic(fmt.Sprintf("Solution to problem %d has not been implemented yet", p))
		case 2:
			panic(err)
		}
	}
}
