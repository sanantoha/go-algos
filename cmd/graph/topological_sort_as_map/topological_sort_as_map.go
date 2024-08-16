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

// O(E + V) time | O(V) space
func sortIter(graph map[string][]string) ([]string, error) {

	cntMap := make(map[string]int, len(graph))

	for k, _ := range graph {
		cntMap[k] = 0
	}

	for _, lst := range graph {
		for _, u := range lst {
			cntMap[u]++
		}
	}

	isCircle := true

	queue := make([]string, 0)

	for k, v := range cntMap {
		if v == 0 {
			queue = append(queue, k)
			isCircle = false
		}
	}

	if isCircle {
		return nil, errors.New("circle in the graph")
	}

	res := make([]string, 0)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		res = append(res, v)

		for _, u := range graph[v] {
			cntMap[u]--

			if cntMap[u] == 0 {
				queue = append(queue, u)
			}
		}
	}

	if len(res) < len(queue) {
		return nil, errors.New("circle in the graph")
	}

	return res, nil
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
		log.Println(err)
	}
	fmt.Println(res)

	res, err = sortIter(graph)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}
