package main

import "fmt"

func numberOfWaysToMakeChange(amount int, denoms []int) int {
	return -1
}

/**
 * https://www.algoexpert.io/questions/Number%20Of%20Ways%20To%20Make%20Change
 * 2 // 1x1  + 1x5 and 6x1
 */
func main() {

	input := []int{1, 5}
	actual := numberOfWaysToMakeChange(6, input)
	fmt.Println(actual)
	fmt.Println(actual == 2)
}
