package ans

import (
	"m9ple/euler/ans/p1"
	"m9ple/euler/ans/p2"
	"m9ple/euler/ans/p3"
	"m9ple/euler/ans/p4"
	"m9ple/euler/ans/p5"
	"m9ple/euler/ans/p6"
)

var solList = map[int]func(){
	1: p1.Sol, 2: p2.Sol, 3: p3.Sol, 4: p4.Sol, 5: p5.Sol, 6: p6.Sol,
}

var askList = map[int]func(){
	1: p1.Ask, 2: p2.Ask, 3: p3.Ask, 4: p4.Ask, 5: p5.Ask, 6: p6.Ask,
}
