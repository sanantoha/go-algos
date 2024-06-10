package main

import (
	"fmt"
)

// O(n) time | O(n) space
func rob(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		dp[i] = max(dp[i-2]+arr[i], dp[i-1])
	}
	return dp[len(arr)-1]
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

func main() {

	nums := []int{4, 1, 2, 7, 5, 3, 1}

	res := rob(nums)
	fmt.Println(res)
	fmt.Println(res == 14)
}
