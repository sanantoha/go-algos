package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"testing"
)

func TestDfs(t *testing.T) {
	graph := createGraph()
	fmt.Println(graph)

	start := 0

	exp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	funcs := []func(*grph.EdgeWeightedDigraph, int) []int{
		dfsRec,
		dfsIter,
	}

	for i, fn := range funcs {
		t.Run("dfs"+strconv.Itoa(i), func(t *testing.T) {
			res := fn(graph, start)
			fmt.Println(res)

			if !reflect.DeepEqual(exp, res) {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}
}

func createGraph() *grph.EdgeWeightedDigraph {
	graph, err := grph.NewEdgeWeightedDigraph(15)
	if err != nil {
		log.Fatal(err)
	}

	edge01, _ := grph.NewDirectedEdge(0, 1, 0)
	edge12, _ := grph.NewDirectedEdge(1, 2, 0)
	edge23, _ := grph.NewDirectedEdge(2, 3, 0)
	edge34, _ := grph.NewDirectedEdge(3, 4, 0)
	edge35, _ := grph.NewDirectedEdge(3, 5, 0)
	edge16, _ := grph.NewDirectedEdge(1, 6, 0)
	edge07, _ := grph.NewDirectedEdge(0, 7, 0)
	edge78, _ := grph.NewDirectedEdge(7, 8, 0)

	edge89, _ := grph.NewDirectedEdge(8, 9, 0)
	edge910, _ := grph.NewDirectedEdge(9, 10, 0)
	edge811, _ := grph.NewDirectedEdge(8, 11, 0)
	edge1112, _ := grph.NewDirectedEdge(11, 12, 0)
	edge1213, _ := grph.NewDirectedEdge(12, 13, 0)
	edge1314, _ := grph.NewDirectedEdge(13, 14, 0)

	graph.AddEdge(edge01)
	graph.AddEdge(edge07)
	graph.AddEdge(edge12)
	graph.AddEdge(edge16)
	graph.AddEdge(edge23)
	graph.AddEdge(edge34)
	graph.AddEdge(edge35)
	graph.AddEdge(edge78)
	graph.AddEdge(edge89)
	graph.AddEdge(edge811)
	graph.AddEdge(edge910)
	graph.AddEdge(edge1112)
	graph.AddEdge(edge1213)
	graph.AddEdge(edge1314)
	return graph
}
