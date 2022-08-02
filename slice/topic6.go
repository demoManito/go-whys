package slice

import "log"

type topic6 struct{}

// A. true
// B. false
// C. panic
// D. 编译失败
func (*topic6) t6() {
	var a [100]int
	var t interface{} = [...]int{99: 0}

	log.Println(a == t)
}
