package mergeSort

import "fmt"

func Merge(arr []int, lo int, mid int, hi int) {
	leftSize := mid - lo
	rightSize := hi - mid + 1
	fmt.Println("leftSize", lo, mid, hi, leftSize)
	fmt.Println("rightSize", lo, mid, hi, rightSize)
	left := make([]int, leftSize)
	right := make([]int, rightSize)
	i := 0
	// 1. 填充left子数组
	for i = lo; i < mid; i += 1 {
		// left[i-lo] = arr[i]
		left[i] = arr[i]
	}
	// 2. 填充right子数组
	for i = mid; i <= hi; i += 1 {
		right[i-mid] = arr[i]
	}
	fmt.Println(left)
	fmt.Println(right)
	// 合并到原数组
	i, j, k := 0, 0, lo
	for i < leftSize && j < rightSize {
		// 如果 left 小于 right
		if left[i] < right[j] {
			// 如果左边小就把
			arr[k] = left[i]
			i++
			k++
		} else {
			arr[k] = right[j]
			j++
			k++
		}
	}
	// 如果循环完了 i还小于right size的话就把全部的i填上去
	for i < leftSize {
		arr[k] = left[i]
		i++
		k++
	}
	// 如果循环完了 i还小于right size的话就把全部的i填上去
	for j < rightSize {
		arr[k] = right[j]
		j++
		k++
	}
}

func MergeSort(arr []int, lo int, hi int) {
	if lo == hi {
		return
	}
	fmt.Println("33---", lo, hi)
	mid := (lo + hi) / 2
	fmt.Println("----", mid)
	MergeSort(arr, lo, mid)
	MergeSort(arr, mid+1, hi)
	Merge(arr, lo, mid-1, hi)
}
