# Miao's Answer of Slice_Exercise

1. B <font color = "red">X C</font>
2. 为什么不是[], [2] <font color = "red">X A</font>
3. C
4. B
5. D
6. D <font color = "red">X C</font>  
     <font color = "deeppink">Question: 是不是因为a原本cap为2，append为创建了新的a，此时传入函数的a就没有再被动过，所以依旧是1，2。。？</font>
7. D
8. D <font color = "red">X B</font>
   <font color = "cyan">'a', 'b', 'c'为对应ascii码，即97，98，99，因此共100个元素</font>
9. A  
    <font color = "deeppink">没懂这是在问什么。。我的理解是：a初始化100个元素均为0，t为第100个元素为0，因此a符合t。。？</font>
10. B
11. 让我选会选D。。但是为什么不是[1, 2, 1]
    <font color = "deeppink">Question: new := old，new是创建了一个新的数组，并在其上slice，而非直接在old对应的数组上是嘛。（但是还是不理解为什么会有4）</font>
12. C  
    <font color = "deeppink">Question: 问题。。应该同上。（感觉不清不楚的。。。</font>
13. A  
    <font color = "deeppink">Question: 函数返回的*old，对应的原数组是否还是1，2，3，4。。？（即：变成了len=2，cap=4的slice）还是说重新创建了一个只有1，2的数组。？</font>