package main

import (
	"reflect"
	"testing"
)

func TestAStarAlgorithm(t *testing.T) {

	startRow := 0
	startCol := 1
	endRow := 4
	endCol := 3

	graph := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}

	expected := [][]int{
		{0, 1},
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{3, 1},
		{4, 1},
		{4, 2},
		{4, 3},
	}

	actual := aStarAlgorithm(startRow, startCol, endRow, endCol, graph)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
