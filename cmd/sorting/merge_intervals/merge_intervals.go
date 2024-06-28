package main

import "fmt"

func merge(intervals [][]int) [][]int {
	return nil
}

func main() {

	intervals := [][]int{
		{1, 5},
		{3, 7},
		{4, 6},
		{6, 8},
	}
	fmt.Println(merge(intervals))

	intervals1 := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(intervals1)

	intervals2 := [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(intervals2)
}
