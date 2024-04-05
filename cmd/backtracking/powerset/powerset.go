package main

import "fmt"

// O(n * 2 ^ n) time | O(n * 2 ^ n) space
func powerset(arr []int) [][]int {

	res := make([][]int, 0)
	res = append(res, make([]int, 0))

	for _, num := range arr {
		subRes := make([][]int, 0)
		for _, lst := range res {
			nlst := make([]int, len(lst))
			copy(nlst, lst)
			nlst = append(nlst, num)
			subRes = append(subRes, nlst)
		}

		for _, sr := range subRes {
			res = append(res, sr)
		}
	}

	return res
}

// O(n * 2 ^ n) time | O(n * 2 ^ n) space
func powersetRec(arr []int) [][]int {
	return helper(arr, len(arr)-1)
}

func helper(arr []int, idx int) [][]int {
	if idx == -1 {
		res := make([][]int, 0)
		res = append(res, make([]int, 0))
		return res
	}

	subRes := helper(arr, idx-1)
	num := arr[idx]
	l := len(subRes)

	for i := 0; i < l; i++ {
		lst := subRes[i]
		nlst := make([]int, len(lst))
		copy(nlst, lst)
		nlst = append(nlst, num)
		subRes = append(subRes, nlst)
	}
	return subRes
}

func main() {
	arr := []int{1, 2, 3}

	fmt.Println(powerset(arr))

	fmt.Println(powersetRec(arr))
}
