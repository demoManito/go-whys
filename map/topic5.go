package _map

import (
	"log"
	"net/url"
)

type topic5 struct{}

type Query struct {
	url.Values // Values type: map[string][]string
}

// A. panic
// B. empty
// C. 编译失败
// D. jesse
func (t *topic5) t1() {
	q := Query{}
	q.Set("name", "jesse")
	log.Println(q.Get("name"))
}
