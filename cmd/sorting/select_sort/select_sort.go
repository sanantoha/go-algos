package main

import (
	"fmt"
	"math/rand"
)

// O(n ^ 2) time | O(1) space
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		if i != minIdx {
			swap(arr, minIdx, i)
		}
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	selectSort(arr)

	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			panic(fmt.Sprintf("array is not sorted: %d != %d", arr[i], arr[i+1]))
		}
	}

	fmt.Println("done")
}
