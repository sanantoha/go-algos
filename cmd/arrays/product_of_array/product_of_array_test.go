package main

import (
	"reflect"
	"testing"
)

func TestProductOfArray(t *testing.T) {
	testTable := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "all integer greater than zero",
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "one zero",
			input:    []int{1, 2, 0, 4},
			expected: []int{0, 0, 8, 0},
		},
		{
			name:     "two zeros",
			input:    []int{0, 2, 0, 4},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := productOfArray(subtest.input)

			if !reflect.DeepEqual(subtest.expected, res) {
				t.Errorf("expected (%v), but got (%v)", subtest.expected, res)
			}
		})
	}
}
