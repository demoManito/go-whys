## ANSWER AND QUESTIONS OF GOROUTINE

1. A  
*`sync.WaitGroup` 作为参数传递给函数时，必须传递指针，否则会复制一份，导致 `WaitGroup` 的计数器不一致，从而导致 panic*