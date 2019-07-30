package mergeSort

import (
	"fmt"
	"testing"
)

// go test -v
func TestMergeSort(t *testing.T) {
	testArr := []int{6, 8, 9, 10, 4, 5, 2, 7}
	fmt.Println(testArr)
	lo := 0
	hi := len(testArr) - 1
	MergeSort(testArr, lo, hi)
	t.Log(testArr)
}

// go test -v
func TestListNodeMergeSort(t *testing.T) {
	a := &ListNode{10, nil}
	for i := 0; i <= 10; i++ {
		t := &ListNode{i, nil}
		t.Next = a
		a = t
	}
	for h := a; h.Next != nil; h = h.Next {
		fmt.Println(h.Val)
	}
	fmt.Println("--------")
	res := ListNodeMergeSort(a)
	for h := res; h.Next != nil; h = h.Next {
		fmt.Println(h.Val)
	}
}
