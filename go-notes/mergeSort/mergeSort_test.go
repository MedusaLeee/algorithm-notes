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
