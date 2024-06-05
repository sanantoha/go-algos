package main

import "fmt"

func lds(arr []int) int {
	return -1
}

func lds1(arr []int) int {
	return -1
}

func ldsList(arr []int) []int {
	return nil
}

func ldsList1(arr []int) []int {
	return nil
}

func main() {

	arr := []int{5, 6, 7, 6, 5, 4, 3, 10, 14, 12} // 5

	fmt.Println(lds(arr))
	fmt.Println(lds1(arr))
	fmt.Println(ldsList(arr))
	fmt.Println(ldsList1(arr))

	arr1 := []int{100, 10, 9, 8, 7, 6, 5, 90, 80, 70, 60, 50, 40, 30, 20} // 9

	fmt.Println(lds(arr1))
	fmt.Println(lds1(arr1))
	fmt.Println(ldsList(arr1))
	fmt.Println(ldsList1(arr1))
}
