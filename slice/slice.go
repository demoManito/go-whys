package slice

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
	println(a)

	a = append(a, 2)
	println(a)
}

// Topic2 .
// A. [1], [2]
// B. [1], [1]
// C. [ ], [1]
// D. panic
func Topic2() {
	a := make([]int, 0, 2)

	add(a)
	println(a[:1])

	a = append(a, 2)
	println(a[:1])
}

// Topic3 .
// A. [2], [ ]
// B. [2], [1]
// C. [2], [2]
// D. panic
func Topic3() {
	a := make([]int, 0, 2)

	a = append(a, 2)
	println(a)

	add(a)
	println(a)
}

// Topic4 .
// A. [1, 2]
// B. [2, 1]
// C. [2, 2]
// D. [1, 1]
func Topic4() {
	a := []int{1, 2}

	swap(a)
	println(a)
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
	println(a)
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
	println(a)
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
	println(a)
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

	println(len(a))
}

// Topic9
// A. true
// B. false
// C. panic
// D. 编译失败
func Topic9() {
	var a [100]int
	var t interface{} = [...]int{99: 0}

	println(a == t)
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
	println(a)
}

// Topic11 .
// A. [1,2,3,4]
// B. [1,2,3,4,1]
// C. [1,2,1,3]
// D. [1,2,1,4]
func Topic11() {
	old := []int{1, 2, 3, 4}
	addOld(old)
	println(old)
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
	println(old)
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
	println(*old)
}

func replacePtr(old *[]int) {
	temp := *old
	*old = temp[0:2]
}

// Topic14 .
// A. [] [1] [1]
// B. [1] [1] [1]
// C. [] [] [1]
// D. panic
func Topic14() {
	a1 := make([]int, 0, 1)
	a2 := a1

	a2 = append(a2, 1)
	println(a1)
	println(a1[:1])
	println(a2)
}

// Topic15 .
// A. [] [1] [1,2,3]
// B. [] [0,0] [1,2,3]
// C. [] [] [1,2,3]
// D. panic
func Topic15() {
	a1 := make([]int, 0, 2)
	a2 := a1

	a2 = append(a2, 1, 2, 3)
	println(a1)
	println(a1[:2])
	println(a2)
}

// Topic16 .
// A. [] [0,0] [1,2,3]
// B. [9] [9,0] [1,2,3]
// C. [9] [0,0] [1,2,3]
// D. panic
func Topic16() {
	a1 := make([]int, 0, 2)
	a2 := a1

	a1 = append(a1, 9)
	a2 = append(a2, 1, 2, 3)
	println(a1)
	println(a1[:2])
	println(a2)
}

// Topic17 .
// A. [] [0,0] [1,2,3,4]
// B. [] [1,0] [1,2,3,4]
// C. [2] [2,3] [2,3,4]
// D. [1] [1,2] [2,3,4]
func Topic17() {
	a1 := make([]int, 0, 2)
	a2 := a1

	a1 = append(a1, 1)
	a2 = append(a2, 2)
	a2 = append(a2, 3)
	a2 = append(a2, 4)
	println(a1)
	println(a1[:2])
	println(a2)
}
