package main

import (
	"fmt"
	"math"
)

// O(n) time | O(1) space
func firstDuplicateValue(arr []int) int {

	for _, val := range arr {
		absVal := int(math.Abs(float64(val)))
		if arr[absVal-1] < 0 {
			return absVal
		}
		arr[absVal-1] *= -1
	}

	return -1
}

func main() {

	fmt.Println(firstDuplicateValue([]int{2, 1, 3, 4, 5, 6, 2, 7, 8, 9}))

	fmt.Println(firstDuplicateValue([]int{}))
}
