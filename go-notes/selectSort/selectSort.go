package selectSort

//type Comparable interface {
//	CompareTo (Comparable) int
//}
//
//type SortInterface interface {
//	Sort([]Comparable)
//	Less()
//}

// 小到大
func SelectSort(arr []int) []int {
	length := len(arr)
	// 第一层循环
	for i := 0; i < length; i += 1 {
		// 默认第一个最小
		minIndex := i
		for j := i + 1; j < length; j += 1 {
			if arr[j] < arr[minIndex] {
				// 换位置
				t := arr[minIndex]
				arr[minIndex] = arr[j]
				arr[j] = t
				minIndex = j
				// 高级替换方法
				// arr[i], arr[minIndex] = arr[minIndex], arr[i]
			}
		}
	}
	return arr
}
