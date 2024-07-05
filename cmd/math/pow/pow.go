package main

import (
	"fmt"
	"math"
)

// O(log(n)) time | O(log(n)) space
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return pow(1/x, -n)
	} else if n%2 == 0 {
		y := pow(x, n/2)
		return y * y
	} else {
		return x * pow(x, n-1)
	}
}

func main() {
	fmt.Println(pow(4, 2))
	fmt.Println(pow(2, 4))
	fmt.Println(pow(2.0, -2))

	fmt.Println(pow(2.1, 3))
	fmt.Println(math.Pow(2.1, 3))
}
