package slice

import "log"

type topic1 struct{}

// A. hello
// B. world
// C. !
// D. panic
func (*topic1) t1() {
	a := map[int]string{1: "hello", 2: "world", 96: "!"}
	log.Println(a['a'])
}
