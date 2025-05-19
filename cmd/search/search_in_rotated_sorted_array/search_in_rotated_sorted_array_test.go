package main

import "testing"

func TestFindSmallestIdx(t *testing.T) {
	testTable := []struct {
		name  string
		input []int
		exp   int
	}{
		{
			name:  "happy path",
			input: []int{40, 50, 60, 70, 80, 90, 0, 10, 20, 30, 31, 32, 33, 34, 35},
			exp:   6,
		},
		{
			name:  "empty",
			input: []int{},
			exp:   -1,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := findSmallestIdx(subtest.input)

			if subtest.exp != res {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		target int
		exp    int
	}{
		{
			name:   "happy path",
			input:  []int{40, 50, 60, 70, 80, 90, 0, 10, 20, 30, 31, 32, 33, 34, 35},
			target: 90,
			exp:    5,
		},
		{
			name:   "empty",
			input:  []int{},
			target: 10,
			exp:    -1,
		},
	}
	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := search(subtest.input, subtest.target)

			if subtest.exp != res {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}

func TestSearch1(t *testing.T) {
	testTable := []struct {
		name   string
		input  []int
		target int
		exp    int
	}{
		{
			name:   "happy path",
			input:  []int{40, 50, 60, 70, 80, 90, 0, 10, 20, 30, 31, 32, 33, 34, 35},
			target: 90,
			exp:    5,
		},
		{
			name:   "empty",
			input:  []int{},
			target: 10,
			exp:    -1,
		},
	}
	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := search1(subtest.input, subtest.target)

			if subtest.exp != res {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}
