package main

import "fmt"

// O(maxSteps ^ height) time | O(maxSteps ^ height) space
func staircaseTraversalRec(height, maxSteps int) int {
	if height <= 1 {
		return 1
	}

	cnt := 0
	for i := 1; i <= min(height, maxSteps); i++ {
		cnt += staircaseTraversalRec(height-i, maxSteps)
	}
	return cnt
}

// O(maxSteps * height) time | O(height) space
func staircaseTraversalRecMemoization(height, maxSteps int) int {
	memo := make([]int, height+1)
	memo[0] = 1
	return helper(height, maxSteps, memo)
}

func helper(height int, maxSteps int, memo []int) int {
	if memo[height] > 0 {
		return memo[height]
	}

	cnt := 0
	for i := 1; i <= min(height, maxSteps); i++ {
		cnt += helper(height-i, maxSteps, memo)
	}

	memo[height] = cnt
	return cnt
}

func staircaseTraversalIter(height, maxSteps int) int {
	return -1
}

func staircaseTraversalSlidingWindow(height, maxSteps int) int {
	return -1
}

func main() {

	fmt.Println(staircaseTraversalRec(4, 2))
	fmt.Println(staircaseTraversalRec(4, 3))
	fmt.Println("==============================================")
	fmt.Println(staircaseTraversalRecMemoization(4, 2))
	fmt.Println(staircaseTraversalRecMemoization(4, 3))
	fmt.Println("==============================================")
	fmt.Println(staircaseTraversalIter(4, 2))
	fmt.Println(staircaseTraversalIter(4, 3))
	fmt.Println("==============================================")
	fmt.Println(staircaseTraversalSlidingWindow(4, 2))
	fmt.Println(staircaseTraversalSlidingWindow(4, 3))
}
