package main

import "fmt"

// O(n ^ 4) time | O(n ^ 2) space
func fourSum(arr []int, target int) [][4]int {
	res := make([][4]int, 0)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				for z := k + 1; z < len(arr); z++ {
					if arr[i]+arr[j]+arr[z]+arr[k] == target {
						res = append(res, [4]int{arr[i], arr[j], arr[z], arr[k]})
					}
				}
			}
		}
	}

	return res
}

// O(n ^ 3) time | O(n ^ 2) space
func fourSum1(arr []int, target int) [][4]int {

	mp := make(map[int][][2]int)

	res := make([][4]int, 0)

	for i := 1; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			compensateKey := target - arr[i] - arr[j]
			lst, ok := mp[compensateKey]
			if ok {
				for _, p := range lst {
					res = append(res, [4]int{arr[i], arr[j], p[0], p[1]})
				}
			}
		}

		for j := i - 1; j >= 0; j-- {
			key := arr[i] + arr[j]
			lst, ok := mp[key]
			if !ok {
				lst = make([][2]int, 0)
			}
			lst = append(lst, [2]int{arr[i], arr[j]})
			mp[key] = lst
		}
	}

	return res
}

func main() {

	fmt.Println(fourSum([]int{7, 6, 4, -1, 1, 2}, 16))

	fmt.Println(fourSum1([]int{7, 6, 4, -1, 1, 2}, 16))
}
