package main

import "fmt"

// O(n * n!) time | O(n * n!) space
func permute(arr []int) [][]int {
	res := make([][]int, 0)
	backtrack(arr, 0, &res)
	return res
}

func backtrack(arr []int, idx int, res *[][]int) {
	if len(arr) == idx {
		*res = append(*res, copySlice(arr))
		return
	}
	for i := idx; i < len(arr); i++ {
		swap(arr, i, idx)
		backtrack(arr, idx+1, res)
		swap(arr, i, idx)
	}
}

func copySlice(arr []int) []int {
	narr := make([]int, len(arr))
	copy(narr, arr)
	return narr
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {

	arr := []int{1, 2, 3}

	fmt.Println(permute(arr))
}
