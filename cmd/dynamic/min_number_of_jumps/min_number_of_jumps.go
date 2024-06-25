package main

import (
	"fmt"
	"math"
)

// O(n ^ 2) time | O(n) space
func minNumberOfJumps(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	mnj := make([]int, len(arr))
	for i, _ := range mnj {
		mnj[i] = math.MaxInt
	}
	mnj[0] = 0

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if j+arr[j] >= i && mnj[i] > mnj[j]+1 {
				mnj[i] = mnj[j] + 1
			}
		}
	}

	return mnj[len(mnj)-1]
}

// O(n) time | O(1) space
func minNumberOfJumps1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	maxReach := arr[0]
	steps := arr[0]
	jumps := 0

	for i := 1; i < len(arr); i++ {
		maxReach := max(i+arr[i], maxReach)
		steps--
		if steps == 0 {
			steps = maxReach - i
			jumps++
		}
	}
	return jumps + 1
}

func main() {

	input := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}

	actual := minNumberOfJumps(input)
	fmt.Println(actual)
	fmt.Println(actual == 4)

	actual = minNumberOfJumps1(input)
	fmt.Println(actual)
	fmt.Println(actual == 4)
}
