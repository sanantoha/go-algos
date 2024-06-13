package main

import (
	"fmt"
	"reflect"
	"slices"
)

// O(n * c) time | O(n * c) space
func knapsackProblem(items [][]int, capacity int) [][]int {
	if items == nil || len(items) == 0 || capacity < 0 {
		return nil
	}

	kp := make([][]int, len(items)+1)
	kp[0] = make([]int, capacity+1)

	for i := 1; i <= len(items); i++ {
		kp[i] = make([]int, capacity+1)
		for c := 1; c <= capacity; c++ {
			item := items[i-1]
			val := item[0]
			weight := item[1]
			if c < weight {
				kp[i][c] = kp[i-1][c]
			} else {
				kp[i][c] = max(kp[i-1][c], kp[i-1][c-weight]+val)
			}
		}
	}

	res := []int{kp[len(items)][capacity]}
	seq := buildSeq(items, kp)
	return [][]int{
		res,
		seq,
	}
}

func buildSeq(items [][]int, kp [][]int) []int {
	res := make([]int, 0)

	i := len(kp) - 1
	c := len(kp[i]) - 1

	for i > 0 {
		if kp[i][c] == kp[i-1][c] {
			i--
		} else {
			i--
			item := items[i]
			res = append(res, i)
			c -= item[1]
		}

		if c == 0 {
			break
		}
	}

	slices.Reverse(res)
	return res
}

func main() {

	input := [][]int{
		{1, 2}, {4, 3}, {5, 6}, {6, 7},
	}

	expected := [][]int{{10}, {1, 3}}

	res := knapsackProblem(input, 10)
	fmt.Println(res)
	fmt.Println(reflect.DeepEqual(res, expected))
}
