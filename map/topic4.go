package _map

type topic4 struct{}

// A. true true true
// B. true true false
// C. panic
// D. 无法编译
func (t *topic4) t1() {
	// var m1 map[int]interface{}
	// var m2 map[int]interface{}
	// log.Println(m1 == nil)
	// log.Println(m2 == nil)
	// log.Println(m1 == m2)
}
