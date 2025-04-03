package main

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testTable := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "happy path",
			input:    []int{10, 9, 37, 9, 17, 11, 0, 44, 23, 16},
			expected: []int{0, 9, 9, 10, 11, 16, 17, 23, 37, 44},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: []int{},
		},
		{
			name:     "one element",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := mergeSort(subtest.input)

			if !reflect.DeepEqual(subtest.expected, res) {
				t.Errorf("expected (%v), but got (%v)", subtest.expected, res)
			}
		})
	}
}
