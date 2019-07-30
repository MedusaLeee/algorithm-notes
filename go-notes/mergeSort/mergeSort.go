package mergeSort

func Merge(arr []int, lo int, mid int, hi int) {
	leftSize := mid - lo + 1
	rightSize := hi - mid
	left := make([]int, leftSize)
	right := make([]int, rightSize)
	// 1. 填充left子数组
	for i := 0; i < leftSize; i++ {
		// left[i-lo] = arr[i]
		left[i] = arr[lo+i]
	}
	// 2. 填充right子数组
	for j := 0; j < rightSize; j++ {
		right[j] = arr[mid+j+1]
	}
	// 合并到原数组
	i, j, k := 0, 0, lo
	for k <= hi && i < leftSize && j < rightSize {
		// 如果 left 小于 right
		if left[i] < right[j] {
			// 如果左边小就把left i放到左边
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	// 如果循环完了 i还小于right size的话就把全部的i填上去
	for i < leftSize && k <= hi {
		arr[k] = left[i]
		i++
		k++
	}
	// 如果循环完了 i还小于right size的话就把全部的i填上去
	for j < rightSize && k <= hi {
		arr[k] = right[j]
		j++
		k++
	}
}

func MergeSort(arr []int, lo int, hi int) {
	if lo < hi {
		mid := (lo + hi) / 2
		MergeSort(arr, lo, mid)
		MergeSort(arr, mid+1, hi)
		Merge(arr, lo, mid, hi)
	}
}
