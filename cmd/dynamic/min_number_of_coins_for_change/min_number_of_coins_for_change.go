package main

import (
	"fmt"
	"math"
)

// O(d * n) time | O(n) space
func minNumberOfCoinsForChange(n int, denoms []int) int {
	if n < 0 || denoms == nil || len(denoms) == 0 {
		return -1
	}

	mnc := make([]int, n+1)

	for i, _ := range mnc {
		mnc[i] = math.MaxInt
	}
	mnc[0] = 0

	for _, denom := range denoms {
		for amount := 1; amount <= n; amount++ {
			if amount >= denom {
				toChange := mnc[amount-denom]
				if toChange != math.MaxInt {
					toChange++
				}
				mnc[amount] = min(mnc[amount], toChange)
			}
		}
	}

	if mnc[n] == math.MaxInt {
		return -1
	} else {
		return mnc[n]
	}
}

func main() {

	input := []int{1, 5, 10}
	actual := minNumberOfCoinsForChange(7, input)
	fmt.Println(actual)
	fmt.Println(actual == 3)
}
