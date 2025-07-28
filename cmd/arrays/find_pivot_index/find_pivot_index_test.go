package main

import (
	"testing"
)

func TestFindPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		exp  int
	}{
		{
			name: "happy path",
			arr:  []int{1, 7, 3, 6, 5, 6},
			exp:  3,
		},
		{
			name: "no pivot index",
			arr:  []int{1, 2, 3},
			exp:  -1,
		},
		{
			name: "first element pivot",
			arr:  []int{2, 1, -1},
			exp:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findPivotIndex(tt.arr)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
