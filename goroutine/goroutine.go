package slice

import (
	"log"
	"sync"
)

// Topic1 .
// A. panic
// B. wait
// C. "run"
func Topic1() {
	run := func(wg sync.WaitGroup) {
		log.Println("run")
		wg.Done()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go run(wg)
	wg.Wait()
}
