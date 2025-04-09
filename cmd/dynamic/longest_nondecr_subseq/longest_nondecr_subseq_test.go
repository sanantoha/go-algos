package main

import (
	"reflect"
	"testing"
)

func TestLnds(t *testing.T) {

	funcs := []func([]int) int{lnds, lnds1}

	input := []int{-2, -1, 2, 3, 4, 5, 2, 2, 2, 2, 3}
	exp := 8

	testTable := []struct {
		name  string
		input []int
		exp   int
	}{
		{
			name:  "happy path",
			input: input,
			exp:   exp,
		},
	}

	for _, fn := range funcs {
		for _, subtest := range testTable {
			t.Run(subtest.name, func(t *testing.T) {
				res := fn(subtest.input)

				if subtest.exp != res {
					t.Errorf("expected (%v), bug to (%v)", subtest.exp, res)
				}
			})
		}
	}
}

func TestLndsList(t *testing.T) {

	funcs := []func([]int) []int{lndsList, lndsList1}

	input := []int{-2, -1, 2, 3, 4, 5, 2, 2, 2, 2, 3}
	exp := []int{-2, -1, 2, 2, 2, 2, 2, 3}

	testTable := []struct {
		name  string
		input []int
		exp   []int
	}{
		{
			name:  "happy path",
			input: input,
			exp:   exp,
		},
	}

	for _, fn := range funcs {
		for _, subtest := range testTable {
			t.Run(subtest.name, func(t *testing.T) {
				res := fn(subtest.input)

				if !reflect.DeepEqual(subtest.exp, res) {
					t.Errorf("expected (%v), bug to (%v)", subtest.exp, res)
				}
			})
		}
	}
}
