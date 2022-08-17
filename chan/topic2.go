package _chan

import (
	"context"
	"log"
)

type topic2 struct {}

// A. 1
// B. 阻塞
// C. panic
func (*topic2) t1()  {
	ctx, _ := context.WithTimeout(context.Background(), 0)
	<-ctx.Done()
	log.Println(1)
}
