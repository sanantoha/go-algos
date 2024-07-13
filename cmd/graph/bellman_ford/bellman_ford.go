package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
)

func main() {
	graph, err := grph.NewEdgeWeightedDigraph(3)

	if err != nil {
		log.Fatalln(err)
	}

	edge0 := grph.DirectedEdge{
		V:      0,
		W:      1,
		Weight: 5,
	}

	edge1 := grph.DirectedEdge{
		V:      0,
		W:      2,
		Weight: 3,
	}

	edge2 := grph.DirectedEdge{
		V:      1,
		W:      2,
		Weight: 6,
	}

	edge3 := grph.DirectedEdge{
		V:      2,
		W:      0,
		Weight: 3,
	}

	graph.AddEdge(&edge0)
	graph.AddEdge(&edge1)
	graph.AddEdge(&edge2)
	graph.AddEdge(&edge3)

	fmt.Println(graph)
}
