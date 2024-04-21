package main

import "fmt"

// O(n) time | O(n) space
func nextGreaterElement(arr []int) []int {

	stack := make([]int, 0)
	res := make([]int, len(arr))
	for i, _ := range arr {
		res[i] = -1
	}

	for i := 0; i < len(arr)*2-1; i++ {
		idx := i % len(arr)

		for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[idx] {
			topIdx := stack[len(stack)-1]
			res[topIdx] = arr[idx]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, idx)
	}

	return res
}

// O(n) time | O(n) space
func nextGreaterElement1(arr []int) []int {

	stack := make([]int, 0)

	res := make([]int, len(arr))
	for i, _ := range arr {
		res[i] = -1
	}

	for i := len(arr)*2 - 1; i >= 0; i-- {
		idx := i % len(arr)

		for len(stack) > 0 {
			if arr[idx] < stack[len(stack)-1] {
				res[idx] = stack[len(stack)-1]
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		stack = append(stack, arr[idx])
	}
	return res
}

func main() {

	input := []int{2, 5, -3, -4, 6, 7, 2}
	actual := nextGreaterElement(input)
	fmt.Println(actual) // [5, 6, 6, 6, 7, -1, 5]

	actual = nextGreaterElement1(input)
	fmt.Println(actual) // [5, 6, 6, 6, 7, -1, 5]
}
