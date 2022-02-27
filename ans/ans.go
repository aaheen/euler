package ans

import (
	"fmt"
	"strconv"

	"github.com/pkg/browser"
)

// Prints out solution for problem p
func Sol(p int) {
	var fn func() = solList[p]
	fn()
}

// Runs interactive solution for problem p
func Ask(p int) {
	var fn func() = askList[p]
	fn()
}

// Opens problem p in my repo on Github
func Page(p int) {
	var gURL string = "https://github.com/m9ple/euler/tree/main/" + strconv.Itoa(p)
	fmt.Println("Heading to", gURL)
	browser.OpenURL(gURL)
}

// Opens my repo's main page on Github
func Repo() {
	browser.OpenURL("https://github.com/m9ple/euler")
}
