package main

import "fmt"

func maxProfit(arr []int) int {
	return -1
}

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices) == 5)

	prices1 := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(prices1) == 5)
}
