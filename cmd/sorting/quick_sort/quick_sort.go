package main

import (
	"fmt"
	"math/rand"
)

// O(n * log(n)) time | O(log(n)) space
func quickSort(arr []int, l, r int) {
	if l > r {
		return
	}

	p := partition(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

func partition(arr []int, l int, r int) int {
	swap(arr, r, l+rand.Intn(r-l+1))

	j := l

	for i := l; i < r; i++ {
		if arr[i] <= arr[r] {
			swap(arr, i, j)
			j++
		}
	}
	swap(arr, r, j)
	return j
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

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			panic(fmt.Sprintf("array is not sorted: %d != %d", arr[i], arr[i+1]))
		}
	}

	fmt.Println("done")
}
