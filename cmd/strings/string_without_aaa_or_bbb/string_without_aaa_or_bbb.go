package main

import "fmt"

func stringWithoutAAAorBBB(a int, b int) string {
	return ""
}

func main() {

	fmt.Println(stringWithoutAAAorBBB(1, 1))
	fmt.Println(stringWithoutAAAorBBB(3, 3))
	fmt.Println(stringWithoutAAAorBBB(2, 5))

	fmt.Println(stringWithoutAAAorBBB(5, 3))

	fmt.Println(stringWithoutAAAorBBB(3, 3))

	fmt.Println(stringWithoutAAAorBBB(1, 4))
}
