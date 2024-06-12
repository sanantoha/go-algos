package main

import (
	"fmt"
)

// O(n) time | O(1) space
func maxProfit(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	minVal := arr[0]
	maxProfitVal := 0

	for _, price := range arr {
		maxProfitVal = max(maxProfitVal, price-minVal)
		minVal = min(price, minVal)
	}
	return maxProfitVal
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices) == 5)

	prices1 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices1) == 0)
}
