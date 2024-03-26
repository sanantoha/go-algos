package main

import "fmt"

// O(n) time | O(n) space
func subarraySum(arr []int, k int) int {

	mp := make(map[int]int, 0)
	mp[0] = 1
	count := 0
	sum := 0

	for _, num := range arr {
		sum += num

		n, ok := mp[sum-k]
		if ok {
			count += n
		}

		mp[sum] = mp[sum] + 1
	}
	return count
}

func main() {

	fmt.Println(subarraySum([]int{1, 1, 1}, 2)) // 2

	fmt.Println(subarraySum([]int{1, 2, 3}, 3)) // 2

}
