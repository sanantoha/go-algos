package main

import (
	"testing"
)

func TestFindPeak(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   map[int]struct{}
	}{
		{
			name:  "happy path 1",
			input: []int{1, 2, 1, 3, 4, 5, 7, 8, 5, 3, 2, 11, 1},
			exp: map[int]struct{}{
				1:  {},
				7:  {},
				11: {},
			},
		},
		{
			name:  "happy path 2",
			input: []int{1, 2, 3, 2, 1},
			exp: map[int]struct{}{
				2: {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findPeak(tt.input)
			if _, exists := tt.exp[res]; !exists {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
