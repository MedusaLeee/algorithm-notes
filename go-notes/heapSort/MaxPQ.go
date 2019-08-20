package heapSort

import "fmt"

type MaxPQ struct {
	heap []int
	n    int
}

func (pq *MaxPQ) Insert(node int) {
	pq.heap[pq.n] = node
	fmt.Println(pq.heap, pq.n)
	// 将第一个节点和最后一个节点交换
	pq.n += 1
	if pq.n > 1 {
		Swap(pq.heap, 0, pq.n-1)
		// 堆化
		BuildHeap(pq.heap, pq.n)
	}
}

func (pq *MaxPQ) DelMax() {
	if pq.n == 0 {
		return
	}
	// 将第一个节点和最后一个节点交换
	Swap(pq.heap, 0, pq.n-1)
	// 置空最后一个节点
	pq.heap[pq.n-1] = 0
	pq.n -= 1
	fmt.Println("del:", pq.heap, pq.n)
	// 堆化
	BuildHeap(pq.heap, pq.n)
}

func (pq *MaxPQ) Heap() []int {
	return pq.heap
}

func (pq *MaxPQ) Size() int {
	return pq.n
}
