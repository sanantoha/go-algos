package main

import (
	"fmt"
	"math/rand"
)

// O(n * log(n)) time | O(1) space
func heapSort(arr []int) {
	endIdx := len(arr) - 1
	buildHeap(arr, endIdx)

	for endIdx > 0 {
		swap(arr, 0, endIdx)
		endIdx--
		siftDown(0, endIdx, arr)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func buildHeap(arr []int, endIdx int) {
	for i := endIdx / 2; i >= 0; i-- {
		siftDown(i, endIdx, arr)
	}
}

func siftDown(currIdx, endIdx int, arr []int) {
	idx := currIdx

	for idx <= endIdx {
		l := idx*2 + 1
		r := idx*2 + 2
		maxIdx := idx

		if l <= endIdx && arr[l] > arr[maxIdx] {
			maxIdx = l
		}
		if r <= endIdx && arr[r] > arr[maxIdx] {
			maxIdx = r
		}

		if idx != maxIdx {
			swap(arr, idx, maxIdx)
			idx = maxIdx
		} else {
			break
		}
	}
}

func main() {

	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	fmt.Println(arr)

	heapSort(arr)

	fmt.Println(arr)
}
