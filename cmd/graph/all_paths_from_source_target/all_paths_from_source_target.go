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

// O(V ^ 2 * 2 ^ V) time | O(2 ^ V) space
func allPathsSourceTarget(graph [][]int) [][]int {

	path := make([]int, 1)
	path[0] = 0

	stack := make([]Info, 1)
	stack[0] = Info{
		v:    0,
		path: path,
	}

	res := make([][]int, 0)

	for len(stack) > 0 {
		info := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(graph)-1 == info.v {
			res = append(res, info.path)
		}

		for _, u := range graph[info.v] {
			p := make([]int, len(info.path))
			copy(p, info.path)
			p = append(p, u)
			stack = append(stack, Info{u, p})
		}
	}
	return res
}

type Info struct {
	v    int
	path []int
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

	fmt.Println(actualRec) // [[0 1 3] [0 2 3]]
	fmt.Println(actual)    // [[0 2 3] [0 1 3]]

	graph1 := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}

	actualRec = allPathsSourceTargetRec(graph1)
	actual = allPathsSourceTarget(graph1)

	fmt.Println(actualRec) // [[0 4] [0 3 4] [0 1 3 4] [0 1 2 3 4] [0 1 4]]
	fmt.Println(actual)    // [[0 1 4] [0 1 2 3 4] [0 1 3 4] [0 3 4] [0 4]]
}
