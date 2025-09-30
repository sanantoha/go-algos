package main

import (
	"testing"
)

func TestFirstDuplicateValue(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		exp   int
	}{
		{
			name:  "happy path",
			input: []int{2, 1, 3, 4, 5, 6, 2, 7, 8, 9},
			exp:   2,
		},
		{
			name:  "empty input slice",
			input: []int{},
			exp:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := firstDuplicateValue(tt.input)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
