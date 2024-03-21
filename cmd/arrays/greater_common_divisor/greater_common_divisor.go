package main

import "fmt"

// O(log(x + y)) time | O(1) space
func greaterCommonDivisor(x, y int) int {
	for y != 0 {
		tmp := x % y
		x = y
		y = tmp
	}
	return x
}

func main() {

	fmt.Println(greaterCommonDivisor(18, 6) == 6)

	fmt.Println(greaterCommonDivisor(18, 10) == 2)

	fmt.Println(greaterCommonDivisor(17, 11) == 1)

	fmt.Println(greaterCommonDivisor(5, 15) == 5)
}
