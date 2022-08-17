package slice

import (
	"log"
)

type topic8 struct {}

// A. [1,2,3,4]
// B. [1,2,3,4,1]
// C. [1,2,1,3]
// D. [1,2,1,4]
func (t *topic8) t1()  {
	old := []int{1, 2, 3, 4}
	t.add(old)
	log.Println(old)
}

func (*topic8) add(old []int) {
	new := old
	old = append(new[0:2], 1)
}

// A. [1,2]
// B. [1,2,3]
// C. [1,2,3,4]
// D. nil
func (t *topic8) t2()  {
	old := []int{1, 2, 3, 4}
	t.replace(old)
	log.Println(old)
}

func (*topic8) replace(old []int) {
	temp := old[0:2]
	old = nil
	old = temp
}

// A. [1,2]
// B. [1,2,3]
// C. [1,2,3,4]
// D. 编译错误
func (t *topic8) t3()  {
	old := &[]int{1, 2, 3, 4}
	t.replacePtr(old)
	log.Println(*old)
}

func (*topic8) replacePtr(old *[]int) {
	temp := *old
	*old = temp[0:2]
}
