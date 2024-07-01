package main

import (
	"fmt"
)

// O(n) time | O(n) space
func waterArea(heights []int) int {

	maximum := make([]int, len(heights))

	leftMax := 0

	for i, h := range heights {
		maximum[i] = leftMax
		leftMax = max(leftMax, h)
	}

	rightMax := 0

	for i := len(heights) - 1; i >= 0; i-- {
		currHeight := min(maximum[i], rightMax)
		diff := currHeight - heights[i]
		if diff > 0 {
			maximum[i] = diff
		} else {
			maximum[i] = 0
		}

		rightMax = max(rightMax, heights[i])
	}

	sum := 0
	for _, v := range maximum {
		sum += v
	}

	return sum
}

// O(n) time | O(1) space
func waterArea1(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	l := 0
	r := len(heights) - 1
	leftMax := heights[l]
	rightMax := heights[r]
	surfaceArea := 0

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, heights[l])
			surfaceArea += leftMax - heights[l]
		} else {
			r--
			rightMax = max(rightMax, heights[r])
			surfaceArea += rightMax - heights[r]
		}
	}
	return surfaceArea
}

func main() {

	input := []int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3}

	actual := waterArea(input)
	fmt.Println(actual == 48)

	actual = waterArea1(input)
	fmt.Println(actual == 48)
}
