package slice

import (
	"log"
	"sync"
)

type topic1 struct{}

// A. panic
// B. wait
// C. print "run"
func (t *topic1) t1() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go t.run(wg)
	wg.Wait()
}

func (t *topic1) run(wg sync.WaitGroup) {
	log.Println("run")
	wg.Done()
}