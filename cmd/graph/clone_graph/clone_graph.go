package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/graph"
)

// O(V + E) time | O(V) space
func cloneGraph(node *graph.Node) *graph.Node {
	if node == nil {
		return node
	}

	cache := make(map[*graph.Node]*graph.Node, 0)
	cache[node] = &graph.Node{Val: node.Val, Neighbors: make([]*graph.Node, 0)}

	stack := make([]*graph.Node, 1)
	stack[0] = node

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		currCopy := cache[curr]

		for _, child := range curr.Neighbors {
			childCopy, ok := cache[child]
			if !ok {
				childCopy = &graph.Node{Val: child.Val, Neighbors: make([]*graph.Node, 0)}
				cache[child] = childCopy
				stack = append(stack, child)
			}
			currCopy.Neighbors = append(currCopy.Neighbors, childCopy)
		}
	}
	return cache[node]
}

func main() {

	n1 := graph.NewNode(1)
	n2 := graph.NewNode(2)
	n3 := graph.NewNode(3)
	n4 := graph.NewNode(4)

	n1.Neighbors = append(n1.Neighbors, n2)
	n1.Neighbors = append(n1.Neighbors, n4)

	n2.Neighbors = append(n2.Neighbors, n1)
	n2.Neighbors = append(n2.Neighbors, n3)

	n3.Neighbors = append(n3.Neighbors, n2)
	n3.Neighbors = append(n3.Neighbors, n4)

	n4.Neighbors = append(n4.Neighbors, n1)
	n4.Neighbors = append(n4.Neighbors, n3)

	res := cloneGraph(n1)

	fmt.Println(res)
}
