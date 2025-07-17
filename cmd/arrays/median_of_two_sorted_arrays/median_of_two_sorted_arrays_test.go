package main

import (
	"reflect"
	"runtime"
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {

	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{7, 8, 9, 10, 11, 12}
	exp := 6.5

	funcs := []func([]int, []int) float64{
		medianOfTwoSortedArrays,
		medianOfTwoSortedArrays1,
	}

	for _, fn := range funcs {
		pc := reflect.ValueOf(fn).Pointer()
		fun := runtime.FuncForPC(pc)

		t.Run(fun.Name(), func(t *testing.T) {
			res := fn(arr1, arr2)

			if res != exp {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}
}
