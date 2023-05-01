package slice

import (
	"sync"
)

// Topic1 .
// A. panic
// B. wait
// C. "run"
func Topic1() {
	run := func(wg sync.WaitGroup) {
		println("run")
		wg.Done()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go run(wg)
	wg.Wait()
}
