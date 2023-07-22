package ans

import (
	"eaheen/euler/ans/p1"
	"eaheen/euler/ans/p2"
	"eaheen/euler/ans/p3"
	"eaheen/euler/ans/p4"
	"eaheen/euler/ans/p45"
	"eaheen/euler/ans/p5"
	"eaheen/euler/ans/p6"
	"eaheen/euler/ans/p7"
	"eaheen/euler/ans/p8"
)

var solList = map[int]func(){
	1: p1.Sol, 2: p2.Sol, 3: p3.Sol, 4: p4.Sol, 5: p5.Sol,
	6: p6.Sol, 7: p7.Sol, 8: p8.Sol, 45: p45.Sol,
}

var askList = map[int]func(){
	1: p1.Ask, 2: p2.Ask, 3: p3.Ask, 4: p4.Ask, 5: p5.Ask,
	6: p6.Ask, 7: p7.Ask, 8: p8.Ask,
}
