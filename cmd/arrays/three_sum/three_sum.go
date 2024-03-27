package main

import (
	"fmt"
	"sort"
)

// O(n ^ 3) time | O(n) space
func threeNumberSum(arr []int, target int) [][3]int {
	res := make([][3]int, 0)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == target {
					res = append(res, [3]int{arr[i], arr[j], arr[k]})
				}
			}
		}
	}
	return res
}

// O(n ^ 2) time | O(n) space
func threeNumberSum1(arr []int, target int) [][3]int {
	sort.Ints(arr)

	res := make([][3]int, 0)

	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i-1] == arr[i] {
			continue
		}

		l := i + 1
		r := len(arr) - 1

		for l < r {
			sum := arr[l] + arr[r] + arr[i]

			if sum == target {
				res = append(res, [3]int{arr[i], arr[l], arr[r]})
				l++
				r--
				for l < r && arr[l-1] == arr[l] {
					l++
				}
				for l < r && arr[r] == arr[r+1] {
					r--
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func main() {
	input := []int{12, 3, 1, 2, -6, 5, -8, 6}

	fmt.Println(threeNumberSum(input, 0))

	fmt.Println(threeNumberSum1(input, 0))
}
