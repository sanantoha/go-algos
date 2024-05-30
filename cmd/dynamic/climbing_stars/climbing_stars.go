package main

import "fmt"

func climbStairsDP(n int) int {
	return -1
}

func climbStairs(n int) int {
	return -1
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
