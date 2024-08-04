package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func dfsRec(graph map[string][]*grph.EdgeT[string], start string) []string {
	return nil
}

func dfsIter(graph map[string][]*grph.EdgeT[string], start string) []string {
	return nil
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
