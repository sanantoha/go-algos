package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

// O(E + V) time | O(V) space
func bfs(graph map[string][]*grph.EdgeT[string], start string) []string {
	if len(graph) == 0 {
		return nil
	}
	queue := make([]string, 1)
	queue[0] = start

	visited := make(map[string]struct{}, 0)
	res := make([]string, 0)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = struct{}{}
		res = append(res, v)

		for _, edge := range graph[v] {
			queue = append(queue, edge.To())
		}

	}

	return res
}

func main() {
	graph, err := grph.NewGraphAsAdjListFromFile("cmd/graph/breadth_first_search/bfs.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(grph.PrintGraphAsAdjList(graph))
	fmt.Println("=====================================")
	// [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24]
	fmt.Println(bfs(graph, "0"))
}
