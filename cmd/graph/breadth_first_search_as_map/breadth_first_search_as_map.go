package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func bfs(graph map[string][]*grph.EdgeT[string], start string) []string {
	return nil
}

func main() {
	graph, err := grph.NewGraphAsAdjListFromFile("cmd/graph/breadth_first_search/bfs.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	fmt.Println(bfs(graph, "0"))
}
