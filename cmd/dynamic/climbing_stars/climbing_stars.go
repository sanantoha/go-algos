package main

import "fmt"

// O(n) time | O(n) space
func climbStairsDP(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

// O(n) time | O(1) space
func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	fst := 1
	snd := 1

	for i := 2; i <= n; i++ {
		fst, snd = snd, fst+snd
	}

	return snd
}

func main() {

	fmt.Println(climbStairsDP(0) == 0)
	fmt.Println(climbStairsDP(1) == 1)
	fmt.Println(climbStairsDP(2) == 2)
	fmt.Println(climbStairsDP(3) == 3)
	fmt.Println(climbStairsDP(4) == 5)
	fmt.Println(climbStairsDP(5) == 8)

	fmt.Println("=======================")

	fmt.Println(climbStairs(0) == 0)
	fmt.Println(climbStairs(1) == 1)
	fmt.Println(climbStairs(2) == 2)
	fmt.Println(climbStairs(3) == 3)
	fmt.Println(climbStairs(4) == 5)
	fmt.Println(climbStairs(5) == 8)
}
