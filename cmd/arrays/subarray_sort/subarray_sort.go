package main

import "fmt"

func subarraySort(arr []int) []int {
	return nil
}

func main() {
	output := subarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})
	fmt.Println(subarraySort(output))

	output1 := subarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 7, 7, 16, 18, 19})
	fmt.Println(subarraySort(output1))

	output2 := subarraySort([]int{1, 2})
	fmt.Println(subarraySort(output2))
}
