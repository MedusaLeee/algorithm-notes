package arraySum

import "fmt"

func ArraySum(s []int, index int) int {
	length := len(s)
	if index == length-1 {
		return s[index]
	}
	return s[index] + ArraySum(s, index+1)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sum := ArraySum(slice, 0)
	fmt.Println(sum)
}
