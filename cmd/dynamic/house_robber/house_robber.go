package main

import (
	"fmt"
)

// O(n) time | O(n) space
func rob(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	fst := arr[0]
	snd := max(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		fst, snd = snd, max(fst+arr[i], snd)
	}

	return snd
}

func main() {

	nums := []int{4, 1, 2, 7, 5, 3, 1}

	res := rob(nums)
	fmt.Println(res)
	fmt.Println(res == 14)
}
