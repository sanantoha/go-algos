package main

import (
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	path := "dijkstraShortestPath.txt"
	graph, err := grph.NewEdgeWeightedDigraphFromFile(path)
	if err != nil {
		t.Fatalf("can not load the graph from path %s", path)
	}

	testTable := []struct {
		name   string
		graph  *grph.EdgeWeightedDigraph
		start  int
		expRes *grph.ShortestPath
	}{
		{
			name:  "happy path",
			graph: graph,
			start: 0,
			expRes: &grph.ShortestPath{
				Shortest: []float64{0, 5, 8, 4, 7},
				Prev:     []int{-1, 3, 1, 0, 3},
			},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			sp, err := findShortestPath(subtest.graph, subtest.start)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(sp.Shortest, subtest.expRes.Shortest) {
				t.Errorf("Shortest expected (%v), but got (%v)", subtest.expRes.Shortest, sp.Shortest)
			}
			if !reflect.DeepEqual(sp.Prev, subtest.expRes.Prev) {
				t.Errorf("Shortest expected (%v), but got (%v)", subtest.expRes.Prev, sp.Prev)
			}
		})
	}
}
