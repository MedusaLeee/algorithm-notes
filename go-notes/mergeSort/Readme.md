# 归并排序


## 习题

### 2.2.4 是否当且仅当两个输入的子数组都有序时原地归并的抽象方法才能得到正确的结果？证明你的结论，或者给出一个反例

必须要两个子数组都有序时归并才能得到正确结果。
如果说数组不有序的话，那么最后只能得到两个数组的混合。
合并后的数组中，属于原有数组的元素的相对顺序不会被改变。
例如子数组 1 3 1 和 2 8 5 原地归并。
结果是 1 2 3 1 8 5，其中 1 3 1 和 2 8 5 的相对顺序不变。

### 2.2.7 证明归并排序的比较次数是单调递增的

比较次数为1/2NlogN--NlgN，因为N >0 lgN为单调递增所以 比较次数为单调递增


### 2.2.17链表排序

[mergeSort_test.go](./mergeSort_test.go)

### 2.2.20间接排序
[mergeSort_test.go](./mergeSort_test.go)

https://github.com/YangXiaoHei/Algorithms/blob/master/Ch_2_2_MergeSort/Practise_2_2_20.java