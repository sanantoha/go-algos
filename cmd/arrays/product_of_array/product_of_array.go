package main

import "fmt"

func productOfArray(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	res := make([]int, len(arr))

	leftProduct := 1

	for i := 0; i < len(arr); i++ {
		res[i] = leftProduct
		leftProduct *= arr[i]
	}

	rightProduct := 1

	for i := len(arr) - 1; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct *= arr[i]
	}

	return res
}

func main() {

	fmt.Println(productOfArray([]int{1, 2, 3, 4}))
	fmt.Println(productOfArray([]int{1, 2, 3, 0}))
	fmt.Println(productOfArray([]int{1, 0, 3, 4}))
	fmt.Println(productOfArray([]int{1, 0, 0, 4}))
}
