# ANSWER AND QUESTIONS OF SLICE

1. B X C
2. [], [2] X A
3. C
4. B
5. D
6. D X C  
    个人理解： a 原本 cap 为 2 ， append 创建了新的 a ，此时传入函数的 a 就没有再被动过，所以依旧是 [1, 2]
7. D
8. D X B  
    笔记：'a', 'b', 'c'为对应ascii码，即97，98，99，因此共100个元素
9. A  
    个人理解： a 初始化 100 个元素均为 0 ， t 为第 100 个元素为 0 ，因此 a 符合 t  
    提问：个人理解是否为题目所表达的意思？
10. B
11. D  
    提问 1 ：为什么不是 [1, 2, 1]  
    个人理解：new := old ， new 创建了一个新的数组，并在其上 slice ，而非直接在 old 对应的数组上  
    提问 2 :为什么输出中含有 4
12.   C  
    提问：同上，问题在于赋值是否创建新数组
13.   A  
    提问：函数返回的 *old ，对应的原数组是否还是 [1，2，3，4] ？ （即：变成了 len=2，cap=4 的 slice ）还是重新创建了一个只有 [1, 2] 的数组?