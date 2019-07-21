package selectSort

import (
	"fmt"
	"testing"
)

// go test -v
func TestSelectSort(t *testing.T) {
	testArr := []int{6, 2, 7, 5, 8, 9}
	fmt.Println()
	res := InsertSort(testArr)
	t.Log(res)
}
