package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

// O(E + V) time | O(V) space
func bfs(graph *grph.EdgeWeightedDigraph, start int) []int {

	visited := make(map[int]struct{}, 0)

	queue := make([]int, 1) // use slice as queue is not optimal, deque takes O(n), better use custom queue impl.
	queue[0] = start
	res := make([]int, 0)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = struct{}{}
		res = append(res, v)

		adj, err := graph.Adj(v)
		if err != nil {
			log.Fatalln(err)
		}
		for _, edge := range adj {
			queue = append(queue, edge.To())
		}
	}

	return res
}

func main() {
	graph, err := grph.NewEdgeWeightedDigraphFromFile("cmd/graph/breadth_search_first/bfs.txt")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(graph)
	log.Println("=====================================")
	log.Println(bfs(graph, 0)) // [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24]
}
