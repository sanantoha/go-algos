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

func buildHeap(arr []int) []int {
	return nil
}

func siftDown(currentIdx int, endIdx int, heap []int) {

}

func siftUp(currentIdx int, heap []int) {

}

func (mh *MinHeap) Peek() int {
	return -1
}

func (mh *MinHeap) Remove() int {
	return -1
}

func (mh *MinHeap) Insert(value int) {

}

func (mh *MinHeap) IsEmpty() bool {
	return len(mh.heap) == 0
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
