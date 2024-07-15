package main

import (
	"container/heap"
	"fmt"
)

// O(n * log(k)) time | O(k) space
func sortKSortedArray(arr []int, k int) []int {

	h := &IntHeap{}

	for i := 0; i <= min(k, len(arr)); i++ {
		heap.Push(h, arr[i])
	}

	idx := 0

	for i := k + 1; i < len(arr); i++ {
		arr[idx] = heap.Pop(h).(int)
		idx++

		heap.Push(h, arr[i])
	}

	for h.Len() > 0 {
		arr[idx] = heap.Pop(h).(int)
		idx++
	}

	return arr
}

// heap impl.
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {

	arr := []int{3, 2, 1, 5, 4, 7, 6, 5}

	fmt.Println(sortKSortedArray(arr, 3))
}
