package heapSort

func Swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func Heapify(tree []int, n int, i int) {
	if i >= n {
		return
	}
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}
	// 相等 就不用排了
	if max != i {
		Swap(tree, max, i)
		Heapify(tree, n, max)
	}
}

func BuildHeap(tree []int, n int) {
	lastNode := n - 1
	var parent int = (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		Heapify(tree, n, i)
	}
}

func HeapSort(tree []int, n int) {
	// 先建一个堆
	BuildHeap(tree, n)
	// 最后一个节点出发
	for i := n - 1; i > 0; i-- {
		// 最后一个节点和第一个节点交换
		Swap(tree, i, 0)
		// 砍断, 相当于这个树的节点已经减少,i 正好代表剩下节点的数量
		Heapify(tree, i, 0)
	}
}
