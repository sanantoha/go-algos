package main

import (
	"fmt"
	"reflect"
)

func diskStacking(disks [][]int) [][]int {
	return nil
}

func main() {

	input := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{2, 2, 8},
		{2, 3, 4},
		{2, 2, 1},
		{4, 4, 5},
	}

	expected := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{4, 4, 5},
	}

	actual := diskStacking(input)
	fmt.Println(actual)

	fmt.Println(reflect.DeepEqual(actual, expected))
}
