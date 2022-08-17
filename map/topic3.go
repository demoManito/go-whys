package slice

type topic3 struct {}

// A. &[1:"one"]
// B. false
// panic
// 无法编译
func (t *topic3) t1() {
	// m := map[int]interface{}{1: "one", 2: "two"}
	// log.Println(&m[1])
}
