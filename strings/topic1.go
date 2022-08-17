package strings

import (
	"log"
	"strings"
)

type topic1 struct {}

// A. \空格trim
// B. trim\空格
// C. trim
// D. rim
func (*topic1) t1() {
	log.Println(strings.Trim("test trim ", "test "))
}
