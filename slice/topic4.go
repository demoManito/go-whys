package slice

import (
	"log"
)

type topic4 struct{}

// A. [1, 2]
// B. [2, 1]
// C. [2, 2]
// D. [1, 1]
func (t *topic4) t4_1() {
	a := []int{1, 2}

	t.swap(a)
	log.Println(a)
}

func (*topic4) swap(a []int) {
	a[0], a[1] = a[1], a[0]
}

// A. [1, 2, 3]
// B. [2, 1, 3]
// C. [1, 2]
// D. [2, 1]
func (t *topic4) t4_2() {
	a := []int{1, 2}

	t.swapAndAdd(a)
	log.Println(a)
}

func (*topic4) swapAndAdd(a []int) {
	a[0], a[1] = a[1], a[0]
	a = append(a, 3)
}

// A. [1, 2, 3]
// B. [2, 1, 3]
// C. [1, 2]
// D. [2, 1]
func (t *topic4) t4_3() {
	a := []int{1, 2}

	t.addAndSwap(a)
	log.Println(a)
}

func (*topic4) addAndSwap(a []int) {
	a = append(a, 3)
	a[0], a[1] = a[1], a[0]
}

// A. [1, 2, 3]
// B. [2, 1, 3]
// C. [1, 2]
// D. [2, 1]
func (t *topic4) t4_4() {
	a := make([]int, 2, 3)
	a[0] = 1
	a[1] = 2

	t.addAndSwap(a)
	log.Println(a)
}
