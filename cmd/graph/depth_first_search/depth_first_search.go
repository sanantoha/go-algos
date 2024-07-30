package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func dfsRec(graph *grph.EdgeWeightedDigraph, start int) []int {
	return nil
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

	log.Println(dfsRec(graph, 0))
	log.Println("=======================================")
	log.Println(dfsIter(graph, 0))
}
