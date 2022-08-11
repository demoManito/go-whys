package slice

import (
	"log"
)

type topic7 struct{}

// A. []
// B. [2]
// C. [1 2]
// D. [1]
func (t *topic7) t7() {
	a := []int{}
	t.add(a)
	a = append(a, 2)
	log.Println(a)
}

func (*topic7) add(a []int) {
	a = append(a, 1)
}
