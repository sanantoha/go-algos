package main

import "fmt"

func search(arr []int, target int) int {
	return -1
}

func search1(arr []int, target int) int {
	return -1
}

func main() {

	arr := []int{40, 50, 60, 70, 80, 90, 0, 10, 20, 30, 31, 32, 33, 34, 35}
	target := 90

	fmt.Println(search(arr, target))
	fmt.Println(search1(arr, target))
}
