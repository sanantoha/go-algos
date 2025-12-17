package main

import "fmt"

type MinHeap struct {
	heap []int
}

func NewMinHeap(arr []int) *MinHeap {
	return &MinHeap{
		heap: buildHeap(arr),
	}
}

// O(n) time | O(n) space
func buildHeap(arr []int) []int {
	heap := make([]int, len(arr))
	copy(heap, arr)
	endIdx := len(heap) - 1
	for i := len(arr) / 2; i >= 0; i-- {
		siftDown(i, endIdx, heap)
	}
	return heap
}

// O(log(n)) time | O(1) space
func siftDown(currentIdx int, endIdx int, heap []int) {

	idx := currentIdx

	for idx <= endIdx {
		l := left(idx)
		r := right(idx)
		minIdx := idx
		if l <= endIdx && heap[l] < heap[minIdx] {
			minIdx = l
		}
		if r <= endIdx && heap[r] < heap[minIdx] {
			minIdx = r
		}
		if minIdx != idx {
			heap[idx], heap[minIdx] = heap[minIdx], heap[idx]
			idx = minIdx
		} else {
			break
		}
	}
}

// O(log(n)) time | O(1) space
func siftUp(currentIdx int, heap []int) {
	idx := currentIdx
	p := parent(idx)
	for p >= 0 && heap[p] > heap[idx] {
		heap[p], heap[idx] = heap[idx], heap[p]
		idx = p
		p = parent(idx)
	}
}

// O(1) time | O(1) space
func (mh *MinHeap) Peek() int {
	if len(mh.heap) == 0 {
		return -1
	}
	return mh.heap[0]
}

// O(log(n)) time | O(1) space
func (mh *MinHeap) Remove() int {
	if len(mh.heap) == 0 {
		return -1
	}

	val := mh.heap[0]
	endIdx := len(mh.heap) - 1
	mh.heap[0], mh.heap[endIdx] = mh.heap[endIdx], mh.heap[0]
	mh.heap = mh.heap[:len(mh.heap)-1]
	endIdx = len(mh.heap) - 1
	siftDown(0, endIdx, mh.heap)
	return val
}

// O(log(n)) time | O(1) space
func (mh *MinHeap) Insert(value int) {
	mh.heap = append(mh.heap, value)
	siftUp(len(mh.heap)-1, mh.heap)
}

func (mh *MinHeap) IsEmpty() bool {
	return len(mh.heap) == 0
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func main() {

	arr := []int{0, 1, 2}

	heap := NewMinHeap(arr)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)

	fmt.Println(heap.heap) // [0, 1, 2, 3, 4, 5]

	fmt.Println(heap.Peek())   // 0
	fmt.Println(heap.Remove()) // 0
	fmt.Println(heap.heap)     // [1, 3, 2, 5, 4]

	fmt.Println(heap.Remove()) // 1
	fmt.Println(heap.Remove()) // 2

	heap.Insert(4)
	fmt.Println(heap.heap) // [3, 4, 4, 5]
}
