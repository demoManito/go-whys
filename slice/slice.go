package slice

import "log"

func add(a []int) {
	a = append(a, 1)
}

// Topic1 .
// A. [1], [2]
// B. [1], [1, 2]
// C. [ ], [2]
// D. [1], [1]
func Topic1() {
	a := make([]int, 0, 2)

	add(a)
	log.Println(a)

	a = append(a, 2)
	log.Println(a)
}

// Topic2 .
// A. [1], [2]
// B. [1], [1]
// C. [ ], [1]
// D. panic
func Topic2() {
	a := make([]int, 0, 2)

	add(a)
	log.Println(a[:1])

	a = append(a, 2)
	log.Println(a[:1])
}

// Topic3 .
// A. [2], [ ]
// B. [2], [1]
// C. [2], [2]
// D. panic
func Topic3() {
	a := make([]int, 0, 2)

	a = append(a, 2)
	log.Println(a)

	add(a)
	log.Println(a)
}

// Topic4 .
// A. [1, 2]
// B. [2, 1]
// C. [2, 2]
// D. [1, 1]
func Topic4() {
	a := []int{1, 2}

	swap(a)
	log.Println(a)
}

func swap(a []int) {
	a[0], a[1] = a[1], a[0]
}

// Topic5 .
// A. [1, 2, 3]
// B. [2, 1, 3]
// C. [1, 2]
// D. [2, 1]
func Topic5() {
	a := []int{1, 2}

	swapAndAdd(a)
	log.Println(a)
}

func swapAndAdd(a []int) {
	a[0], a[1] = a[1], a[0]
	a = append(a, 3)
}

// Topic6 .
// A. [1, 2, 3]
// B. [2, 1, 3]
// C. [1, 2]
// D. [2, 1]
func Topic6() {
	a := []int{1, 2}

	addAndSwap(a)
	log.Println(a)
}

func addAndSwap(a []int) {
	a = append(a, 3)
	a[0], a[1] = a[1], a[0]
}

// Topic7 .
// A. [1, 2, 3]
// B. [2, 1, 3]
// C. [1, 2]
// D. [2, 1]
func Topic7() {
	a := make([]int, 2, 3)
	a[0] = 1
	a[1] = 2

	addAndSwap(a)
	log.Println(a)
}

// Topic8 .
// A. 99
// B. 100
// C. 1000
// D. 3
func Topic8() {
	a := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	a['a'] = 1000

	log.Println(len(a))
}

// Topic9
// A. true
// B. false
// C. panic
// D. 编译失败
func Topic9() {
	var a [100]int
	var t interface{} = [...]int{99: 0}

	log.Println(a == t)
}

// Topic10 .
// A. []
// B. [2]
// C. [1 2]
// D. [1]
func Topic10() {
	a := []int{}
	add(a)
	a = append(a, 2)
	log.Println(a)
}

// Topic11 .
// A. [1,2,3,4]
// B. [1,2,3,4,1]
// C. [1,2,1,3]
// D. [1,2,1,4]
func Topic11() {
	old := []int{1, 2, 3, 4}
	addOld(old)
	log.Println(old)
}

func addOld(old []int) {
	new := old
	old = append(new[0:2], 1)
}

// Topic12 .
// A. [1,2]
// B. [1,2,3]
// C. [1,2,3,4]
// D. nil
func Topic12() {
	old := []int{1, 2, 3, 4}
	replace(old)
	log.Println(old)
}

func replace(old []int) {
	temp := old[0:2]
	old = nil
	old = temp
}

// Topic13
// A. [1,2]
// B. [1,2,3]
// C. [1,2,3,4]
// D. 编译错误
func Topic13() {
	old := &[]int{1, 2, 3, 4}
	replacePtr(old)
	log.Println(*old)
}

func replacePtr(old *[]int) {
	temp := *old
	*old = temp[0:2]
}
