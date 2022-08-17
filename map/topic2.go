package slice

import (
	"log"
	"sync"
)

type topic2 struct {}

// A. [1:1 2:2 3:3 ...]
// B. [0:0 1:1 2:2 3:3 ...]
// C. panic
// D. 无法编译
func (t *topic2) t1() {
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
		log.Println(m)
}
