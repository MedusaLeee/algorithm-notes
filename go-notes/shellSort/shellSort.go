package shellSort

func ShellSort(arr []int) []int {
	n := len(arr)
	h := 1
	for h < n/3 { // 寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		// 将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			// 间隔h插入排序
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				swap(arr, j, j-h)
			}
		}
		h /= 3
	}
	return arr
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
