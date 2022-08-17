package _chan

import "log"

type topic1 struct {}

// A. 1
// B. 阻塞
// C. 编译失败
// D. panic
func (*topic1) t1()  {
	c := make(chan int, 1)
	<-c
	log.Println(1)
}
