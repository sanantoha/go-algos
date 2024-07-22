package main

import "fmt"

// O(V * 2 ^ V) time | O(V) space
func allPathsSourceTargetRec(graph [][]int) [][]int {
	start := 0
	ans := make([]int, 1)
	ans[0] = 0
	res := make([][]int, 0)
	backtrack(graph, start, &ans, &res)
	return res
}

func backtrack(graph [][]int, v int, ans *[]int, res *[][]int) {
	if v == len(graph)-1 {
		nans := make([]int, len(*ans))
		copy(nans, *ans)
		*res = append(*res, nans)
		return
	}

	for _, u := range graph[v] {
		*ans = append(*ans, u)
		backtrack(graph, u, ans, res)
		*ans = (*ans)[:len(*ans)-1]
	}
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

	fmt.Println(actualRec) // [[0,4], [0,3,4], [0,1,3,4], [0,1,2,3,4], [0,1,4]]
	fmt.Println(actual)    // [[0, 4], [0, 3, 4], [0, 1, 3, 4], [0, 1, 2, 3, 4], [0, 1, 4]]
}
