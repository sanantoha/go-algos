package main

import (
	"fmt"
	"math/rand"
)

// O(n * log(n)) time | O(n) space
func mergeSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	mergeSortHelper(arr, 0, len(arr)-1, res)
	return res
}

func mergeSortHelper(arr []int, l int, r int, res []int) {
	if l >= r {
		return
	}

	mid := int(uint(l+r) >> 1)
	mergeSortHelper(res, l, mid, arr)
	mergeSortHelper(res, mid+1, r, arr)
	merge(arr, l, mid, r, res)
}

func merge(arr []int, l int, mid int, r int, res []int) {
	i := l
	j := mid + 1

	idx := l

	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			res[idx] = arr[i]
			idx++
			i++
		} else {
			res[idx] = arr[j]
			idx++
			j++
		}
	}

	for i <= mid {
		res[idx] = arr[i]
		idx++
		i++
	}
	for j <= r {
		res[idx] = arr[j]
		idx++
		j++
	}
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
