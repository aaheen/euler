package ans

import (
	"fmt"

	"github.com/pkg/browser"
)

// Prints out solution for problem p
func Sol(p int) {
	x := 1
	defer recAns(p, &x)
	var fn func() = solMap[p]
	x = 2
	fn()
}

// Runs interactive solution for problem p
func Ask(p int) {
	x := 1
	defer recAns(p, &x)
	var fn func() = askMap[p]
	x = 2
	fn()
}

// Opens writeup for problem p on my personal blog site
// DEPRECATED
//func Writeup(p int) {
//	var gURL string = "https://heen.dev/writeup/euler/" + strconv.Itoa(p)
//	fmt.Println("Heading to", gURL)
//	browser.OpenURL(gURL)
//}

// Opens my repo's main page on Github
func Repo() {
	var gURL string = "https://github.com/aaheen/euler"
	fmt.Println("Heading to", gURL)
	browser.OpenURL(gURL)
}

// Recovers if Sol or Ask fail
func recAns(p int, x *int) {
	if err := recover(); err != nil {
		switch *x {
		case 1:
			panic(fmt.Sprintf("Solution to problem %d has not been implemented yet", p))
		case 2:
			panic(err)
		}
	}
}
