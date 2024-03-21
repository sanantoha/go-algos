package main

import (
	"fmt"
	"math"
)

// O(n ^ 2) time | O(n) space
func minRewards(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	rewards := makeInitRewards(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			rewards[i] = rewards[i-1] + 1
		} else {
			j := i - 1
			for j >= 0 && arr[j] > arr[j+1] {
				rewards[j] = int(math.Max(float64(rewards[j]), float64(rewards[j+1]+1)))
				j--
			}
		}
	}

	return rewardsSum(rewards)
}

// O(n) time | O(n) space
func minRewards1(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	rewards := makeInitRewards(arr)

	for _, i := range getMinIdx(arr) {
		extendsIdxs(arr, rewards, i)
	}

	return rewardsSum(rewards)
}

func extendsIdxs(arr []int, rewards []int, idx int) {

	l := idx - 1
	for l >= 0 && arr[l] > arr[l+1] {
		rewards[l] = int(math.Max(float64(rewards[l]), float64(rewards[l+1])+1))
		l--
	}

	r := idx + 1
	for r < len(arr) && arr[r-1] < arr[r] {
		rewards[r] = rewards[r-1] + 1
		r++
	}
}

func getMinIdx(arr []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(arr); i++ {

		if i == 0 && arr[i] < arr[i+1] {
			res = append(res, i)
		} else if i == len(arr)-1 && arr[i-1] > arr[i] {
			res = append(res, i)
		} else if i == 0 || i == len(arr)-1 {
			continue
		} else if arr[i-1] > arr[i] && arr[i] < arr[i+1] {
			res = append(res, i)
		}
	}

	return res
}

// O(n) time | O(n) space
func minRewards2(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	rewards := makeInitRewards(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			rewards[i] = rewards[i-1] + 1
		}
	}

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			rewards[i] = int(math.Max(float64(rewards[i]), float64(rewards[i+1]+1)))
		}
	}

	return rewardsSum(rewards)
}

func makeInitRewards(arr []int) []int {
	rewards := make([]int, len(arr))
	for i, _ := range arr {
		rewards[i] = 1
	}
	return rewards
}

func rewardsSum(rewards []int) int {
	sum := 0
	for _, v := range rewards {
		sum += v
	}
	return sum
}

func main() {
	arr := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}

	fmt.Println(minRewards(arr))
	fmt.Println(minRewards1(arr))
	fmt.Println(minRewards2(arr))

}
