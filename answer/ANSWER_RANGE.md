## ANSWER AND QUESTIONS OF RANGE
1. A  
*users 为两个指向结构体 user 的指针组成的数组，第一个循环依次将 users 中指针复制进 usersCopy ，第二个循环打印结构体对应字段的名称和内容*

2. C  
*在循环中，在 usersCopy 中添加的是 &u ，两次迭代都是，因此在结束循环时都指向同一个元素*  
__Q：为什么循环结束后指针不会无效化？__
