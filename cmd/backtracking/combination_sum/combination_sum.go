package main

import "fmt"

// O(N ^ (M / T + 1)) time | O(T / M) space
// where N is the number of candidates, M is the smallest candidate among all the given integers,
// and T is the target value.
// Thus the time complexity is exponential and this is expected because the algorithm is
// recursive backtracking.
func combinationSum(arr []int, target int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	backtrack(arr, target, 0, &ans, &res)
	return res
}

func backtrack(arr []int, target int, idx int, ans *[]int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		nans := make([]int, len(*ans))
		copy(nans, *ans)
		*res = append(*res, nans)
		return
	}

	for i := idx; i < len(arr); i++ {
		*ans = append(*ans, arr[i])
		backtrack(arr, target-arr[i], i, ans, res)
		*ans = (*ans)[:len(*ans)-1]
	}
}

func main() {

	arr := []int{2, 3, 5, 7}

	fmt.Println(combinationSum(arr, 7)) // // [[2, 2, 3], [2, 5], [7]]
}
