package slice

import (
	"fmt"
	"log"
)

// Topic1 .
// A. a b
// B. b a
// C. a a
// D. b b
func Topic1() {
	f := func() { log.Println("a") }
	defer f()
	f = func() { log.Println("b") }
	defer f()
}

// Topic2 .
// A. 1 2 3 4 5 6 7 8 9 10
// B. 1 2 3 4 5 6 7 8 9
// C. 9 9 9 9 9 9 9 9 9 9
// D. 10 10 10 10 10 10 10 10 10 10
func Topic2() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

// Topic3 .
// A. 11
// B. 01
// C. 10
// D. 00
func Topic3() {
	incr := func(i int) int {
		i++
		return i
	}

	i := 0

	defer fmt.Print(incr(i))
	fmt.Print(i)
}

// Topic4 .
// A. 11
// B. 01
// C. 10
// D. 00
func Topic4() {
	i := 0

	incr := func() int {
		i++
		return i
	}

	defer fmt.Print(incr())
	fmt.Print(i)
}
