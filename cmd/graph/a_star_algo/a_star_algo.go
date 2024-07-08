package main

import (
	"fmt"
	"reflect"
)

func aStarAlgorithm(startRow int, startCol int, endRow int, endCol int, graph [][]int) [][]int {
	return nil
}

func main() {

	startRow := 0
	startCol := 1
	endRow := 4
	endCol := 3

	graph := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}

	expected := [][]int{
		{0, 1},
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{3, 1},
		{4, 1},
		{4, 2},
		{4, 3},
	}

	actual := aStarAlgorithm(startRow, startCol, endRow, endCol, graph)
	fmt.Println(actual)

	fmt.Println(reflect.DeepEqual(expected, actual))
}
