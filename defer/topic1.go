package slice

import "log"

type topic1 struct{}

// A. a b
// B. b a
// C. a a
// D. b b
func (*topic1) t1() {
	f := func() { log.Println("a") }
	defer f()
	f = func() { log.Println("b") }
	defer f()
}
