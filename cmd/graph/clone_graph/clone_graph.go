package main

import (
	"fmt"
	"github.com/sanantoha/go-algos/internals/graph"
)

func cloneGraph(node *graph.Node) *graph.Node {
	return nil
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
