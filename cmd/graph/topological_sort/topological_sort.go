package main

import (
	"errors"
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

// O(E + V) time | O(V) space
func sortRec(graph *grph.Digraph) ([]int, error) {

	visited := make([]int, graph.V)
	stack := make([]int, 0)

	for v := 0; v < graph.V; v++ {
		if visited[v] == 0 {
			err := dfs(graph, visited, v, &stack)
			if err != nil {
				return nil, err
			}
		}
	}

	res := make([]int, graph.V)
	idx := 0
	for len(stack) > 0 {
		res[idx] = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		idx++
	}
	return res, nil
}

func dfs(graph *grph.Digraph, visited []int, v int, stack *[]int) error {
	visited[v] = 1

	adj, err := graph.Adj(v)
	if err != nil {
		return err
	}
	for _, u := range adj {
		if visited[u] == 1 {
			return errors.New(fmt.Sprintf("graph has circle in %d", v))
		}
		if visited[u] == 0 {
			err = dfs(graph, visited, u, stack)
		}
	}

	visited[v] = 2
	*stack = append(*stack, v)
	return nil
}

func sortIter(graph *grph.Digraph) ([]int, error) {
	return nil, nil
}

func main() {

	graph, err := grph.NewDigraphFromFile("cmd/graph/topological_sort/digraph.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(graph)

	// [8 9 1 0 2 4 3 5 6 7 10 11 12 13]
	res, err := sortRec(graph)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)

	// [0, 1, 8, 2, 9, 4, 3, 5, 6, 10, 7, 11, 12, 13]
	res, err = sortIter(graph)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
