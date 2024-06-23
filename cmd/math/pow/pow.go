package main

import (
	"fmt"
	"math"
)

func pow(x float64, n int) float64 {
	return -1
}

func main() {
	fmt.Println(pow(4, 2))
	fmt.Println(pow(2, 4))
	fmt.Println(pow(2.0, -2))

	fmt.Println(pow(2.1, 3))
	fmt.Println(math.Pow(2.1, 3))
}
