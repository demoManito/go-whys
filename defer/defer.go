package slice

// Topic1 .
// A. a b
// B. b a
// C. a a
// D. b b
func Topic1() {
	f := func() { println("a") }
	defer f()
	f = func() { println("b") }
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
			println(i)
		}()
	}
}

// Topic3 .
// A. 1 1
// B. 0 1
// C. 1 0
// D. 0 0
func Topic3() {
	incr := func(i int) int {
		i++
		return i
	}

	i := 0

	defer println(incr(i))
	println(i)
}

// Topic4 .
// A. 1 1
// B. 0 1
// C. 1 0
// D. 0 0
func Topic4() {
	i := 0

	incr := func() int {
		i++
		return i
	}

	defer println(incr())
	println(i)
}
