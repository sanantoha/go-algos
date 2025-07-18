package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiskStacking(t *testing.T) {
	input := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{2, 2, 8},
		{2, 3, 4},
		{2, 2, 1},
		{4, 4, 5},
	}

	expected := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{4, 4, 5},
	}

	actual := diskStacking(input)
	fmt.Println(actual)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
