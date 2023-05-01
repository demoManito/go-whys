package _map

import (
	"net/url"
	"sync"
)

// Topic1
// A. hello
// B. world
// C. !
// D. panic
func Topic1() {
	a := map[int]string{1: "hello", 2: "world", 96: "!"}
	println(a['a'])
}

// Topic2
// A. [1:1 2:2 3:3 ...]
// B. [0:0 1:1 2:2 3:3 ...]
// C. panic
// D. 无法编译
func Topic2() {
	m := make(map[int]interface{})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			m[i] = i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			delete(m, i)
		}
	}()
	wg.Wait()
	println(m)
}

// Topic3 .
// A. &[1:"one"]
// B. false
// panic
// 无法编译
func Topic3() {
	// m := map[int]interface{}{1: "one", 2: "two"}
	// println(&m[1])
}

// Topic4 .
// A. true true true
// B. true true false
// C. panic
// D. 无法编译
func Topic4() {
	// var m1 map[int]interface{}
	// var m2 map[int]interface{}
	// println(m1 == nil)
	// println(m2 == nil)
	// println(m1 == m2)
}

// Topic5 .
// A. panic
// B. empty
// C. 编译失败
// D. jesse
func Topic5() {
	var query struct {
		url.Values // Values type: map[string][]string
	}
	query.Set("name", "jesse")
	println(query.Get("name"))
}
