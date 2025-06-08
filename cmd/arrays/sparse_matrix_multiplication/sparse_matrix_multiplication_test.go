package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestSparseMatrixMultiplication(t *testing.T) {

	mat1 := [][]int{
		{1, 0, 0},
		{-1, 0, 3},
	}

	mat2 := [][]int{
		{7, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}

	exp := [][]int{
		{7, 0, 0},
		{-7, 0, 3},
	}

	funcs := []func([][]int, [][]int) [][]int{
		multiply,
		multiply1,
	}

	for i, fn := range funcs {
		t.Run("multiply"+strconv.Itoa(i), func(t *testing.T) {
			res := fn(mat1, mat2)
			fmt.Println(res)

			if !reflect.DeepEqual(exp, res) {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}
}
