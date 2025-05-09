package main

import (
	"reflect"
	"testing"
)

var arr0 = []int{1, 2, 3, 6, -100, -90, -80, -70, -60, 7, 8, 9, 10, -50, -40}
var exp0 = 9
var expList0 = []int{-100, -90, -80, -70, -60, 7, 8, 9, 10}

var arr = []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
var exp = 6
var expList = []int{10, 22, 33, 50, 60, 80}

var arr1 = []int{4, 10, 4, 3, 8, 9}
var exp1 = 3

var arr2 = []int{10, 9, 2, 5, 3, 7, 101, 18}
var exp2 = 4

var arr3 = []int{1, -10, 20, 30, 2, 3, 4, 5}
var exp3 = 5

func TestLis(t *testing.T) {
	testTable := []struct {
		name  string
		fn    func([]int) int
		input []int
		exp   int
	}{
		{
			name:  "happy path lis0 = 9",
			fn:    lis0,
			input: arr0,
			exp:   exp0,
		},
		{
			name:  "happy path lis0 = 6",
			fn:    lis0,
			input: arr,
			exp:   exp,
		},
		{
			name:  "happy path lis0 = 3",
			fn:    lis0,
			input: arr1,
			exp:   exp1,
		},
		{
			name:  "happy path lis0 = 4",
			fn:    lis0,
			input: arr2,
			exp:   exp2,
		},
		{
			name:  "happy path lis0 = 5",
			fn:    lis0,
			input: arr3,
			exp:   exp3,
		},
		{
			name:  "happy path lis = 9",
			fn:    lis,
			input: arr0,
			exp:   exp0,
		},
		{
			name:  "happy path lis = 6",
			fn:    lis,
			input: arr,
			exp:   exp,
		},
		{
			name:  "happy path lis = 3",
			fn:    lis,
			input: arr1,
			exp:   exp1,
		},
		{
			name:  "happy path lis = 4",
			fn:    lis,
			input: arr2,
			exp:   exp2,
		},
		{
			name:  "happy path lis = 5",
			fn:    lis,
			input: arr3,
			exp:   exp3,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := subtest.fn(subtest.input)

			if subtest.exp != res {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}

func TestLisList(t *testing.T) {
	testTable := []struct {
		name  string
		fn    func([]int) []int
		input []int
		exp   []int
	}{
		{
			name:  "happy path lis0 = 9",
			fn:    lisList0,
			input: arr0,
			exp:   expList0,
		},
		{
			name:  "happy path lis0 = 6",
			fn:    lisList0,
			input: arr,
			exp:   []int{10, 22, 33, 50, 60, 80},
		},
		{
			name:  "happy path lis0 = 3",
			fn:    lisList0,
			input: arr1,
			exp:   []int{4, 8, 9},
		},
		{
			name:  "happy path lis0 = 4",
			fn:    lisList0,
			input: arr2,
			exp:   []int{2, 5, 7, 101},
		},
		{
			name:  "happy path lis0 = 5",
			fn:    lisList0,
			input: arr3,
			exp:   []int{1, 2, 3, 4, 5},
		},
		{
			name:  "happy path lis = 9",
			fn:    lisList,
			input: arr0,
			exp:   expList0,
		},
		{
			name:  "happy path lis = 6",
			fn:    lisList,
			input: arr,
			exp:   []int{10, 22, 33, 41, 60, 80},
		},
		{
			name:  "happy path lis = 3",
			fn:    lisList,
			input: arr1,
			exp:   []int{3, 8, 9},
		},
		{
			name:  "happy path lis = 4",
			fn:    lisList,
			input: arr2,
			exp:   []int{2, 3, 7, 18},
		},
		{
			name:  "happy path lis = 5",
			fn:    lisList,
			input: arr3,
			exp:   []int{-10, 2, 3, 4, 5},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := subtest.fn(subtest.input)

			if !reflect.DeepEqual(subtest.exp, res) {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}
