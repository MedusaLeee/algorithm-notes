package selectSort

// 小到大
func InsertSort(arr []int) []int {
	length := len(arr)
	// 第一层循环
	for i := 0; i < length; i += 1 {
		// 默认第一个最小
		for j := i; j > 0 && arr[j] < arr[j-1]; j -= 1 {
			t := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = t
		}
	}
	return arr
}
