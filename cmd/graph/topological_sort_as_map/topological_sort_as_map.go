package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// O(E + V) time | O(V) space
func sort(graph map[string][]string) ([]string, error) {

	visited := make(map[string]int, 0)
	stack := make([]string, 0)
	res := make([]string, 0)

	for v, _ := range graph {
		if visited[v] == 0 {
			err := dfs(graph, visited, v, &stack)
			if err != nil {
				return nil, err
			}
		}
	}

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, v)
	}
	return res, nil
}

func dfs(graph map[string][]string, visited map[string]int, v string, stack *[]string) error {
	visited[v] = 1

	for _, u := range graph[v] {
		if visited[u] == 1 {
			return errors.New("circle in the graph " + v)
		}
		if visited[u] == 0 {
			err := dfs(graph, visited, u, stack)
			if err != nil {
				return err
			}
		}
	}

	visited[v] = 2
	*stack = append(*stack, v)
	return nil
}

func sortIter(graph map[string][]string) ([]string, error) {
	return nil, nil
}

func main() {
	graph := make(map[string][]string, 0)

	graph["A"] = []string{"B", "C", "D"}
	graph["B"] = []string{"C"}
	graph["C"] = []string{"D"}
	//graph["D"] = []string{"A", "B"}
	graph["D"] = []string{}

	res, err := sort(graph)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)

	res, err = sortIter(graph)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
