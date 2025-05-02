package main

import (
	"fmt"
	grph "github.com/sanantoha/go-algos/internals/graph"
	"reflect"
	"testing"
)

func TestSortRec(t *testing.T) {
	graph, err := createGraph(t)

	fmt.Println(graph)

	exp := []int{8, 9, 1, 0, 2, 4, 3, 5, 6, 7, 10, 11, 12, 13}
	res, err := sortRec(graph)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, err)
	}
}

func TestSortIter(t *testing.T) {
	graph, err := createGraph(t)

	fmt.Println(graph)

	exp := []int{0, 1, 8, 2, 9, 3, 4, 5, 10, 6, 11, 7, 12, 13}
	res, err := sortIter(graph)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, err)
	}
}

func createGraph(t *testing.T) (*grph.Digraph, error) {
	V := 14
	graph, err := grph.NewDigraph(V)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	_ = graph.AddEdge(0, 2)
	_ = graph.AddEdge(1, 3)
	_ = graph.AddEdge(2, 3)
	_ = graph.AddEdge(2, 4)
	_ = graph.AddEdge(3, 5)
	_ = graph.AddEdge(4, 5)
	_ = graph.AddEdge(5, 10)
	_ = graph.AddEdge(5, 6)
	_ = graph.AddEdge(6, 7)
	_ = graph.AddEdge(7, 12)
	_ = graph.AddEdge(8, 9)
	_ = graph.AddEdge(9, 10)
	_ = graph.AddEdge(10, 11)
	_ = graph.AddEdge(11, 12)
	_ = graph.AddEdge(12, 13)
	return graph, err
}
