package main

import (
	"fmt"
	"math"
)

// O(log(n)) time | O(1) space
func sqrt(n int) int {
	if n <= 0 {
		return -1
	}

	var l uint = 0
	var r uint = uint(n)

	for l <= r {
		mid := (l + r) >> 1

		if mid*mid <= uint(n) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return int(l) - 1
}

func main() {
	fmt.Println(sqrt(4))
	fmt.Println(sqrt(5))
	fmt.Println(sqrt(6))
	fmt.Println(sqrt(7))
	fmt.Println(sqrt(8))
	fmt.Println(sqrt(9))
	fmt.Println(sqrt(16))

	fmt.Println(math.Sqrt(7))
}
