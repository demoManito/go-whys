package slice

import (
	"log"
)

type topic1 struct{}

// A. [1], [2]
// B. [1], [1, 2]
// C. [ ], [2]
// D. [1], [1]
func (t *topic1) t1() {
	a := make([]int, 0, 2)

	t.add(a)
	log.Println(a)

	a = append(a, 2)
	log.Println(a)
}

func (*topic1) add(a []int) {
	a = append(a, 1)
}
