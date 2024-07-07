package main

import "fmt"

func allPathsSourceTargetRec(graph [][]int) [][]int {
	return nil
}

func allPathsSourceTarget(graph [][]int) [][]int {
	return nil
}

func main() {

	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}

	actualRec := allPathsSourceTargetRec(graph)
	actual := allPathsSourceTarget(graph)

	fmt.Println(actualRec) // [[0, 1, 3], [0, 2, 3]]
	fmt.Println(actual)    // [[0, 1, 3], [0, 2, 3]]

	graph1 := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}

	actualRec = allPathsSourceTargetRec(graph1)
	actual = allPathsSourceTarget(graph1)

	fmt.Println(actualRec) // [[0, 4], [0, 3, 4], [0, 1, 3, 4], [0, 1, 2, 3, 4], [0, 1, 4]]
	fmt.Println(actual)    // [[0,4], [0,3,4], [0,1,3,4], [0,1,2,3,4], [0,1,4]]
}
