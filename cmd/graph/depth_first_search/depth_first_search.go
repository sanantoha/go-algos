package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

// O(E + V) time | O(V) space
func dfsRec(graph *grph.EdgeWeightedDigraph, start int) []int {

	res := make([]int, 0)
	visited := make([]bool, graph.V)

	dfs(graph, visited, start, &res)

	return res
}

func dfs(graph *grph.EdgeWeightedDigraph, visited []bool, v int, res *[]int) {
	if visited[v] {
		return
	}
	visited[v] = true
	*res = append(*res, v)

	adj, err := graph.Adj(v)
	if err != nil {
		log.Fatalln(err)
	}
	for _, edge := range adj {
		dfs(graph, visited, edge.To(), res)
	}
}

func dfsIter(graph *grph.EdgeWeightedDigraph, start int) []int {
	return nil
}

func main() {
	graph, err := grph.NewEdgeWeightedDigraphFromFile("cmd/graph/depth_first_search/dfs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(graph)

	log.Println("=======================================")
	log.Println(dfsRec(graph, 0))
	log.Println("=======================================")
	log.Println(dfsIter(graph, 0))
}
