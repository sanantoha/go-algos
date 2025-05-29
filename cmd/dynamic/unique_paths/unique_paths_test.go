package main

import "testing"

func TestUniquePath(t *testing.T) {
	tests := []struct {
		name   string
		m      int
		n      int
		expRes int
	}{
		{
			name:   "happy path 3 x 2",
			m:      3,
			n:      2,
			expRes: 3,
		},
		{
			name:   "happy path 3 x 7",
			m:      3,
			n:      7,
			expRes: 28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := uniquePaths(tt.m, tt.n)

			if tt.expRes != res {
				t.Errorf("expected %v, but got %v", tt.expRes, res)
			}
		})
	}
}
