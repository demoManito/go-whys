## ANSWER AND QUESTIONS OF DEFER

1. B  
*defer 将函数压入栈中，先进后出*

2. C  
*defer 了 10 次 println(i) 语句，而在循环结束时 i 为 9 ，因此打印 10 次 9*

3. B  
*defer 了一个闭包，即传入 i=0 后不再干扰外界；后续打印 i=0 ，再看 defer 的闭包，打印 i=1*

1. A  
*由于 i 为全局变量， defer 的内容不构成闭包，其中的 i 与外界的 i 统一，因此incr 先将 i 自增 1 后，外界再进行打印，因此打印两个 1*