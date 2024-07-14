package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func sortRec(graph *grph.Digraph) []int {
	return nil
}

func sortIter(graph *grph.Digraph) []int {
	return nil
}

func main() {

	graph, err := grph.NewDigraphFromFile("cmd/graph/topological_sort/digraph.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(graph)

	// [8, 9, 1, 0, 2, 3, 4, 5, 10, 11, 6, 7, 12, 13]
	res := sortRec(graph)
	fmt.Println(res)

	// [0, 1, 8, 2, 9, 4, 3, 5, 6, 10, 7, 11, 12, 13]
	res = sortIter(graph)
	fmt.Println(res)
}
