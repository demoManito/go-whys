package strings

import (
	"strings"
)

// Topic1 .
// A. \空格trim
// B. trim\空格
// C. trim
// D. rim
func Topic1() {
	println(strings.Trim("test trim ", "test "))
}
