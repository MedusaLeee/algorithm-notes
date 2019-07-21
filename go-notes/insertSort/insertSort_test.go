package selectSort

import (
	"fmt"
	"testing"
)

// go test -v
func TestInsertSort(t *testing.T) {
	testArr := []int{6, 2, 7, 5, 8, 9}
	fmt.Println()
	res := InsertSort(testArr)
	t.Log(res)
}
