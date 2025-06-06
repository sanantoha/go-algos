package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKnapsackProblem(t *testing.T) {

	input := [][]int{
		{1, 2}, {4, 3}, {5, 6}, {6, 7},
	}

	expected := [][]int{{10}, {1, 3}}

	res := knapsackProblem(input, 10)
	fmt.Println(res)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v, but got %v", expected, res)
	}
}
