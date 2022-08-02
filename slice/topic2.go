package slice

import (
	"log"
)

type topic2 struct{}

// A. [1], [2]
// B. [1], [1]
// C. [ ], [1]
// D. panic
func (t *topic2) t2() {
	a := make([]int, 0, 2)

	t.add(a)
	log.Println(a[:1])

	a = append(a, 2)
	log.Println(a[:1])
}

func (*topic2) add(a []int) {
	a = append(a, 1)
}
