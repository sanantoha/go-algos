package main

import (
	"strconv"
	"testing"
)

func TestCanFinish(t *testing.T) {

	funcs := []func(int, [][]int) bool{
		canFinish,
		canFinish1,
	}

	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		exp           bool
	}{
		{
			name:          "case1",
			numCourses:    2,
			prerequisites: [][]int{},
			exp:           true,
		},
		{
			name:       "case2",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
			exp: true,
		},
		{
			name:       "case3",
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
				{0, 1},
			},
			exp: false,
		},
		{
			name:       "case4",
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0},
				{2, 1},
				{3, 2},
				{0, 3},
			},
			exp: false,
		},
		{
			name:       "case5",
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0},
				{2, 1},
				{3, 2},
			},
			exp: true,
		},
	}

	for i, fn := range funcs {
		for _, tt := range tests {
			t.Run(tt.name+" func "+strconv.Itoa(i), func(t *testing.T) {
				res := fn(tt.numCourses, tt.prerequisites)

				if res != tt.exp {
					t.Errorf("expected %v, but to %v", tt.exp, res)
				}
			})
		}
	}
}
