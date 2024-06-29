package main

import (
	"fmt"
	"math/rand"
)

func mergeSort(arr []int) []int {
	return arr
}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	arr = mergeSort(arr)

	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			panic(fmt.Sprintf("array is not sorted: %d != %d", arr[i], arr[i+1]))
		}
	}

	fmt.Println("done")
}
