package cardSort

import (
	"fmt"
	"testing"
)

// go test -v
func TestCardSort(t *testing.T) {
	testArr := []int{6, 2, 7, 5, 8, 9}
	fmt.Println()
	cardSort := &CardSort{testArr}
	res := cardSort.Sort()
	t.Log(res)
}
