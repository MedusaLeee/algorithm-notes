package binarySearch

import "fmt"

func binarySearch(s []int, k int) int {
	low := 0
	high := len(s) - 1
	for low < high {

		var mid int = (low + high) / 2
		guess := s[mid]
		if guess == k {
			return mid
		}
		if guess > k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	result := binarySearch(arr, 3)
	fmt.Println(result)
	result = binarySearch(arr, 4)
	fmt.Println(result)
	result = binarySearch(arr, 6)
	fmt.Println(result)
}
