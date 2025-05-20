package main

import (
	"reflect"
	"testing"
)

var graph = [][]int{
	{1, 2},
	{3},
	{3},
	{},
}

var graph1 = [][]int{
	{4, 3, 1},
	{3, 2, 4},
	{3},
	{4},
	{},
}

func TestAllPathsSourceTargetRec(t *testing.T) {
	testTable := []struct {
		name  string
		input [][]int
		exp   [][]int
	}{
		{
			name:  "happy path graph",
			input: graph,
			exp: [][]int{
				{0, 1, 3},
				{0, 2, 3},
			},
		},
		{
			name:  "happy path graph1",
			input: graph1,
			exp: [][]int{
				{0, 4},
				{0, 3, 4},
				{0, 1, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 4},
			},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := allPathsSourceTargetRec(subtest.input)

			if !reflect.DeepEqual(subtest.exp, res) {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}

func TestAllPathsSourceTarget(t *testing.T) {
	testTable := []struct {
		name  string
		input [][]int
		exp   [][]int
	}{
		{
			name:  "happy path graph",
			input: graph,
			exp: [][]int{
				{0, 2, 3},
				{0, 1, 3},
			},
		},
		{
			name:  "happy path graph1",
			input: graph1,
			exp: [][]int{
				{0, 1, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 3, 4},
				{0, 3, 4},
				{0, 4},
			},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := allPathsSourceTarget(subtest.input)

			if !reflect.DeepEqual(subtest.exp, res) {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}
