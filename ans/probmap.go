package ans

import (
	"m9ple/euler/ans/p1"
	"m9ple/euler/ans/p2"
	"m9ple/euler/ans/p3"
	"m9ple/euler/ans/p4"
)

var solList = map[int]func(){
	1: p1.Sol, 2: p2.Sol, 3: p3.Sol, 4: p4.Sol,
}

var askList = map[int]func(){
	1: p1.Ask, 2: p2.Ask, 3: p3.Ask, 4: p4.Ask,
}
