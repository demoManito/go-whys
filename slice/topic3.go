package slice

import (
	"log"
)

type topic3 struct{}

// A. [2], [ ]
// B. [2], [1]
// C. [2], [2]
// D. panic
func (t *topic3) t3() {
	a := make([]int, 0, 2)

	a = append(a, 2)
	log.Println(a)

	t.add(a)
	log.Println(a)
}

func (*topic3) add(a []int) {
	a = append(a, 1)
}
