/*
	Though it may seem eccentric, this is (in spirit) the best way I've found
	to implement this. There doesn't seem to be an easy way to evaluate a
	string value into a function call as I had hoped, I will update this with
	a better solution if I happen to find one in my adventures.
*/

package ans

import (
	"aaheen/euler/ans/p1"
	"aaheen/euler/ans/p2"
	"aaheen/euler/ans/p3"
	"aaheen/euler/ans/p4"
	"aaheen/euler/ans/p45"
	"aaheen/euler/ans/p5"
	"aaheen/euler/ans/p6"
	"aaheen/euler/ans/p7"
	"aaheen/euler/ans/p8"
	"aaheen/euler/ans/p9"
)

// Storing the functions for every problem in an interface
type funcMap map[int]func()

var solMap = funcMap{
	1: p1.Sol, 2: p2.Sol, 3: p3.Sol, 4: p4.Sol, 5: p5.Sol,
	6: p6.Sol, 7: p7.Sol, 8: p8.Sol, 9: p9.Sol, 45: p45.Sol,
}

var askMap = funcMap{
	1: p1.Ask, 2: p2.Ask, 3: p3.Ask, 4: p4.Ask, 5: p5.Ask,
	6: p6.Ask, 7: p7.Ask, 8: p8.Ask, 9: p9.Ask,
}
