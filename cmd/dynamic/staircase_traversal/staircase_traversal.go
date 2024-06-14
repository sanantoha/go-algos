package main

import "fmt"

func staircaseTraversalRec(height, maxSteps int) int {
	return -1
}

func staircaseTraversalRecMemoization(height, maxSteps int) int {
	return -1
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
