package strings

import (
	"log"
	"strings"
)

// Topic1 .
// A. \空格trim
// B. trim\空格
// C. trim
// D. rim
func Topic1() {
	log.Println(strings.Trim("test trim ", "test "))
}
