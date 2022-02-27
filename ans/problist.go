package ans

import (
	"m9ple/euler/ans/p0"
	"m9ple/euler/ans/p1"
	"m9ple/euler/ans/p2"
	"m9ple/euler/ans/p3"
	"m9ple/euler/ans/p4"
)

var solList = []func(){
	p0.Sol, p1.Sol, p2.Sol, p3.Sol, p4.Sol,
}

var askList = []func(){
	p0.Sol, p1.Ask, p2.Ask, p3.Ask, p4.Ask,
}
