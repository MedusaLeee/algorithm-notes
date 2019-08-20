package heapSort

import (
	"testing"
)

// go test -v
func TestHeapSort(t *testing.T) {
	tree := []int{4, 10, 3, 5, 1, 2}
	n := 6
	Heapify(tree, n, 0)
	t.Log(tree)
	tree2 := []int{6, 2, 7, 5, 8, 9}
	n2 := 6
	BuildHeap(tree2, n2)
	t.Log(tree2)
	tree3 := []int{6, 2, 7, 5, 8, 9}
	n3 := 6
	HeapSort(tree3, n3)
	t.Log(tree3)
	maxPQ := &MaxPQ{heap: make([]int, 10, 10), n: 0}
	maxPQ.Insert(3)
	maxPQ.Insert(5)
	maxPQ.Insert(8)
	maxPQ.Insert(1)
	maxPQ.Insert(6)
	maxPQ.Insert(9)
	t.Log("size1:", maxPQ.Size())
	t.Log("heap1:", maxPQ.Heap())
	maxPQ.DelMax()
	t.Log("size2:", maxPQ.Size())
	t.Log("heap2:", maxPQ.Heap())
}
