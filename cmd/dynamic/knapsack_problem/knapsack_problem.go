package main

import (
	"fmt"
	"reflect"
)

func knapsackProblem(items [][]int, capacity int) [][]int {

	return nil
}

func main() {

	input := [][]int{
		{1, 2}, {4, 3}, {5, 6}, {6, 7},
	}

	expected := [][]int{{10}, {1, 3}}

	res := knapsackProblem(input, 10)
	fmt.Println(res)
	fmt.Println(reflect.DeepEqual(res, expected))
}
