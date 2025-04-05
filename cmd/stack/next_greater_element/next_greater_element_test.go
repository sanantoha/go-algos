package main

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	testTable := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "happy path",
			input:    []int{2, 5, -3, -4, 6, 7, 2},
			expected: []int{5, 6, 6, 6, 7, -1, 5},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := nextGreaterElement(subtest.input)

			if !reflect.DeepEqual(subtest.expected, res) {
				t.Errorf("expected (%v), but got (%v)", subtest.expected, res)
			}
		})
	}
}

func TestNextGreaterElement1(t *testing.T) {
	testTable := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "happy path",
			input:    []int{2, 5, -3, -4, 6, 7, 2},
			expected: []int{5, 6, 6, 6, 7, -1, 5},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := nextGreaterElement1(subtest.input)

			if !reflect.DeepEqual(subtest.expected, res) {
				t.Errorf("expected (%v), but got (%v)", subtest.expected, res)
			}
		})
	}
}
