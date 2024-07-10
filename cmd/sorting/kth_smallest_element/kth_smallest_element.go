package main

import (
	"fmt"
	"math"
	"math/rand"
)

func kthSmallestElement(arr []int, k int) int {
	return helper(arr, 0, len(arr)-1, k)
}

func helper(arr []int, l int, r int, k int) int {
	if l > r {
		return math.MinInt
	}

	p := partition(arr, l, r)
	if p == k-1 {
		return arr[p]
	}
	if p < k-1 {
		return helper(arr, p+1, r, k)
	}
	return helper(arr, l, p-1, k)
}

func partition(arr []int, l int, r int) int {
	p := l + rand.Intn(r-l+1)
	swap(arr, p, r)

	j := l
	for i := l; i < r; i++ {
		if arr[i] <= arr[r] {
			swap(arr, i, j)
			j++
		}
	}

	swap(arr, j, r)
	return j
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := []int{8, 2, 1, 6, 9, 3, 45, 6, 7, 13}

	for i := 0; i < len(arr); i++ {
		fmt.Println(kthSmallestElement(arr, i+1))
	}

	fmt.Println(kthSmallestElement(arr, -1))
	fmt.Println(kthSmallestElement(arr, 0))
	fmt.Println(kthSmallestElement(arr, 11))
}
