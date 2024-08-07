package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"slices"
)

// O(E + V) time | O(V) space
func dfsRec(graph map[string][]*grph.EdgeT[string], start string) []string {
	res := make([]string, 0)
	visited := make(map[string]bool, 0)
	helper(graph, visited, start, &res)
	return res
}

func helper(graph map[string][]*grph.EdgeT[string], visited map[string]bool, v string, res *[]string) {
	if visited[v] {
		return
	}
	visited[v] = true
	*res = append(*res, v)

	for _, edge := range graph[v] {
		helper(graph, visited, edge.To(), res)
	}
}

// O(E + V) time | O(V) space
func dfsIter(graph map[string][]*grph.EdgeT[string], start string) []string {

	stack := make([]string, 1)
	stack[0] = start
	visited := make(map[string]bool, 0)
	res := make([]string, 0)

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[v] {
			continue
		}
		visited[v] = true
		res = append(res, v)

		edges := graph[v]
		slices.Reverse(edges)
		for _, edge := range edges {
			stack = append(stack, edge.To())
		}
	}
	return res
}

func main() {
	graph, err := grph.NewGraphAsAdjListFromFile("cmd/graph/depth_first_search/dfs.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	fmt.Println(dfsRec(graph, "0"))
	fmt.Println("=====================================")
	fmt.Println(dfsIter(graph, "0"))
}
