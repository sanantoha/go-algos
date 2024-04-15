package main

import (
	"fmt"
	"math"
)

// O(1) time | O(1) space
func reverse(n int) int {

	var res int64 = 0

	for n != 0 {
		res = res*10 + int64(n)%10
		if res < math.MinInt || res > math.MaxInt {
			return 0
		}
		n /= 10
	}
	return int(res)
}

func main() {

	fmt.Println(reverse(123))

	fmt.Println(reverse(0))

	fmt.Println(reverse(-123))

	fmt.Println(reverse(120))
}
