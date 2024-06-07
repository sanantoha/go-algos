package main

import "fmt"

func lnds(arr []int) int {
	return -1
}

func lnds1(arr []int) int {
	return -1
}

func lndsList(arr []int) []int {
	return nil
}

func lndsList1(arr []int) []int {
	return nil
}

func main() {

	arr := []int{-2, -1, 2, 3, 4, 5, 2, 2, 2, 2, 3} // 8
	// arr := []int{-1, 3, 4, 5, 2, 2, 2, 2} // 5

	fmt.Println(lnds(arr))
	fmt.Println(lnds1(arr))
	fmt.Println(lndsList(arr))
	fmt.Println(lndsList1(arr))
}
