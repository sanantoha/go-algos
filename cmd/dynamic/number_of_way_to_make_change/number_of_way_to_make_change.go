package main

import "fmt"

// O(n * d) time | O(n) space
func numberOfWaysToMakeChange(n int, denoms []int) int {
	if n < 0 || denoms == nil || len(denoms) == 0 {
		return -1
	}

	nwmc := make([]int, n+1)
	nwmc[0] = 1

	for _, denom := range denoms {
		for amount := 1; amount <= n; amount++ {
			if amount >= denom {
				nwmc[amount] += nwmc[amount-denom]
			}
		}
	}

	return nwmc[n]
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
