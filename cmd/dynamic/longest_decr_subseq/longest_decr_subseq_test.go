package main

import (
	"reflect"
	"testing"
)

var input = []int{5, 6, 7, 6, 5, 4, 3, 10, 14, 12}

var exp = 5

var input2 = []int{100, 10, 9, 8, 7, 6, 5, 90, 80, 70, 60, 50, 40, 30, 20}

var exp2 = 9

func TestLds(t *testing.T) {

	fns := []func([]int) int{
		lds,
		lds1,
	}

	tests := []struct {
		name  string
		input []int
		exp   int
	}{
		{
			name:  "lds 5",
			input: input,
			exp:   exp,
		},
		{
			name:  "lds 9",
			input: input2,
			exp:   exp2,
		},
	}

	for _, tt := range tests {
		for _, fn := range fns {
			t.Run(tt.name, func(t *testing.T) {
				res := fn(tt.input)

				if res != tt.exp {
					t.Errorf("expected %v, but got %v", tt.exp, res)
				}
			})
		}
	}
}

func TestLdsList(t *testing.T) {
	fns := []func([]int) []int{
		ldsList,
		ldsList1,
	}

	tests := []struct {
		name  string
		input []int
		exp   []int
	}{
		{
			name:  "ldsList 5",
			input: input,
			exp:   []int{7, 6, 5, 4, 3},
		},
		{
			name:  "ldsList 9",
			input: input2,
			exp:   []int{100, 90, 80, 70, 60, 50, 40, 30, 20},
		},
	}

	for _, tt := range tests {
		for _, fn := range fns {
			t.Run(tt.name, func(t *testing.T) {

				res := fn(tt.input)

				if !reflect.DeepEqual(res, tt.exp) {
					t.Errorf("expected %v, but got %v", tt.exp, res)
				}
			})
		}
	}
}
