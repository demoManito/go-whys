package _chan

import (
	"context"
	"log"
)

// Topic1 .
// A. 1
// B. 阻塞
// C. 编译失败
// D. panic
func Topic1() {
	c := make(chan int, 1)
	<-c
	log.Println(1)
}

// Topic2 .
// A. 1
// B. 阻塞
// C. panic
func Topic2() {
	ctx, _ := context.WithTimeout(context.Background(), 0)
	<-ctx.Done()
	log.Println(1)
}
