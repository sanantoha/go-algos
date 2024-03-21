package main

import "fmt"

// O(n) time | O(1) space
func findPivotIndex(arr []int) int {

	leftSum := 0
	rightSum := 0
	for _, v := range arr {
		rightSum += v
	}

	for i := 0; i < len(arr); i++ {
		num := arr[i]
		rightSum -= num

		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}
	return -1
}

func main() {
	fmt.Println(findPivotIndex([]int{1, 7, 3, 6, 5, 6}) == 3)

	fmt.Println(findPivotIndex([]int{1, 2, 3}) == -1)

	fmt.Println(findPivotIndex([]int{2, 1, -1}) == 0)
}
