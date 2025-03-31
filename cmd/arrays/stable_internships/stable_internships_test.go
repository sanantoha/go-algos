package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestStableInternships(t *testing.T) {
	testTable := []struct {
		name    string
		interns [][]int
		teams   [][]int
		expRes  [][]int
	}{
		{
			name: "happy path",
			interns: [][]int{
				{0, 1, 2},
				{0, 2, 1},
				{1, 2, 0},
			},
			teams: [][]int{
				{2, 1, 0},
				{0, 1, 2},
				{0, 1, 2},
			},
			expRes: [][]int{
				{0, 1},
				{1, 0},
				{2, 2},
			},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := stableInternships(subtest.interns, subtest.teams)

			sort.Slice(res, func(i, j int) bool {
				return res[i][0] < res[j][0]
			})
			if !reflect.DeepEqual(subtest.expRes, res) {
				t.Errorf("expected (%v), but got (%v)", subtest.expRes, res)
			}
		})
	}
}
