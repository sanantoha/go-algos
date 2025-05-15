package main

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testTable := []struct {
		name  string
		input []int
		exp   []int
	}{
		{
			name:  "happy path",
			input: []int{27, 40, 9, 18, 26, 28, 37, 26, 29, 34},
			exp:   []int{9, 18, 26, 26, 27, 28, 29, 34, 37, 40},
		},
		{
			name:  "empty slice",
			input: []int{},
			exp:   []int{},
		},
		{
			name:  "nil",
			input: nil,
			exp:   nil,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			heapSort(subtest.input)

			if !reflect.DeepEqual(subtest.exp, subtest.input) {
				t.Errorf("expected %v, but got %v", subtest.exp, subtest.input)
			}
		})
	}
}
