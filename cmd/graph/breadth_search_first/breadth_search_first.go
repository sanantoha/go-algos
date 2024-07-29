package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func bfs(graph *grph.EdgeWeightedDigraph, start int) []int {
	return nil
}

func main() {
	graph, err := grph.NewEdgeWeightedDigraphFromFile("cmd/graph/breadth_search_first/bfs.txt")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(graph)
	log.Println("=====================================")
	log.Println(bfs(graph, 0))
}
