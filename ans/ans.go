package ans

import (
	"fmt"
	"strconv"

	"github.com/pkg/browser"
)

// Prints out solution for problem p
func Sol(p int) {
	defer recP(p)
	var fn func() = solList[p]
	fn()
}

// Runs interactive solution for problem p
func Ask(p int) {
	defer recP(p)
	var fn func() = askList[p]
	fn()
}

// Recover if Sol, Ask, or solution code panics
func recP(p int) {
	err := recover()
	switch err {
	case nil:
		return
	default:
		fmt.Println("Error trying to run problem", p, ":\n", err)
	}
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
