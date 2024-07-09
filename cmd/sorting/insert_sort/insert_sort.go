package main

import (
	"fmt"
	"math/rand"
)

// O(n ^ 2) time | O(1) space
func insertSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	insertSort(arr)

	fmt.Println(arr)
}
