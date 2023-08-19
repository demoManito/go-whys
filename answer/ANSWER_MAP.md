## ANSWER AND QUESTIONS OF MAP

1. C  
*'a' 实际为 96*

2. C  
*两个并发同时对 wg 进行写入与删除，导致 panic*

3. D  
*map 元素无法取地址*

4. D  
*map 只能与 nil 进行比较，无法与另一个 map 进行比较*

5. A  
*query.url.Values 并未初始化 Values 这个 map ，为 nil ，尝试调用为 nil 的 map 中的元素会导致 panic*
