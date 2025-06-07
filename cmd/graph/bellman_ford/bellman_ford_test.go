package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	expShortest := []float64{-9.0, -20.0, -18.0, -2.0, -11.0}
	expPrev := []int{4, 2, 4, 0, 1}
	exp := &grph.ShortestPath{
		Shortest: expShortest,
		Prev:     expPrev,
	}
	expCircle := []int{1, 2, 4}

	graph, err := grph.NewEdgeWeightedDigraph(5)
	if err != nil {
		t.Fatal(err)
	}

	edge, _ := grph.NewDirectedEdge(0, 1, 6)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(0, 3, 7)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(1, 2, 5)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(2, 1, -2)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(1, 3, 8)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(1, 4, -4)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(3, 2, -3)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(4, 2, -7)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(4, 0, 2)
	graph.AddEdge(edge)
	edge, _ = grph.NewDirectedEdge(3, 4, 9)
	graph.AddEdge(edge)

	fmt.Println(graph)

	sp := findShortestPath(graph, 0)
	if !reflect.DeepEqual(exp, sp) {
		t.Errorf("expected %v, but got %v", exp, sp)
	}

	circle := findNegativeWeightCycle(graph, sp)
	if !reflect.DeepEqual(expCircle, circle) {
		t.Errorf("expected %v, but got %v", expCircle, circle)
	}
}
