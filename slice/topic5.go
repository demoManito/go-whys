package slice

import "log"

type topic5 struct{}

// A. 99
// B. 100
// C. 1000
// D. 3
func (*topic5) t5() {
	a := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	a['a'] = 1000

	log.Println(len(a))
}
