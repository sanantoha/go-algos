package main

import "fmt"

func lis0(arr []int) int {
	return -1
}

func lis(arr []int) int {
	return -1
}

func lisList0(arr []int) []int {
	return nil
}

func lisList(arr []int) []int {
	return nil
}

func main() {

	arr0 := []int{1, 2, 3, 6, -100, -90, -80, -70, -60, 7, 8, 9, 10, -50, -40} // 9
	fmt.Println(lis0(arr0))
	fmt.Println(lis(arr0))
	fmt.Println(lisList0(arr0))
	fmt.Println(lisList(arr0))
	fmt.Println("==============================")

	arr := []int{10, 22, 9, 33, 21, 50, 41, 60, 80} // 6
	fmt.Println(lis0(arr))
	fmt.Println(lis(arr))
	fmt.Println(lisList0(arr))
	fmt.Println(lisList(arr))
	fmt.Println("==============================")

	arr1 := []int{4, 10, 4, 3, 8, 9} // 3
	fmt.Println(lis0(arr1))
	fmt.Println(lis(arr1))
	fmt.Println(lisList0(arr1))
	fmt.Println(lisList(arr1))
	fmt.Println("==============================")

	arr2 := []int{10, 9, 2, 5, 3, 7, 101, 18} // 4
	fmt.Println(lis0(arr2))
	fmt.Println(lis(arr2))
	fmt.Println(lisList0(arr2))
	fmt.Println(lisList(arr2))
	fmt.Println("==============================")

	arr3 := []int{1, -10, 20, 30, 2, 3, 4, 5} // 5
	fmt.Println(lis0(arr3))
	fmt.Println(lis(arr3))
	fmt.Println(lisList0(arr3))
	fmt.Println(lisList(arr3))
}
