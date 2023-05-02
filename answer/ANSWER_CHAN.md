## ANSWER AND QUESTIONS OF CHAN

1. D  
*channel 没有发送方，导致死锁 panic*

2. A  
*`context.WithTimeout` timeout 为 0 时，会立即返回，不会阻塞*

3. D  
*无缓冲的 channel 是同步的，必须有接收才能发送，否则会阻塞。这里接收方是 main goroutine 所以会阻塞，导致 panic*

