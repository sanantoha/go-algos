package main

import "testing"

func TestWordSearch(t *testing.T) {

	input := [][]rune{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	tests := []struct {
		name   string
		input  [][]rune
		word   string
		expRes bool
	}{
		{
			name: "happy path",
			input: [][]rune{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:   "ABCCE",
			expRes: true,
		},
		{
			name:   "happy path 1",
			input:  input,
			word:   "ABCESEEEFS",
			expRes: true,
		},
		{
			name:   "happy path 2",
			input:  input,
			word:   "ABCEFSADEESE",
			expRes: true,
		},
		{
			name:   "happy path 3",
			input:  input,
			word:   "ABCEV",
			expRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := exist(tt.input, tt.word)

			if tt.expRes != res {
				t.Errorf("expected %v, but got %v", tt.expRes, res)
			}
		})
	}
}
