package main

import (
	"fmt"
	"math/rand"
)

// O(n) time | O(n) space
func countingSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}

	minVal := arr[0]
	maxVal := arr[0]

	for _, v := range arr {
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}

	cs := make([]int, maxVal-minVal+1)
	for _, v := range arr {
		cs[v-minVal]++
	}

	idx := 0
	for i := 0; i < len(cs); i++ {
		for j := 0; j < cs[i]; j++ {
			arr[idx] = i + minVal
			idx++
		}
	}
}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	countingSort(arr)

	fmt.Println(arr)
}
