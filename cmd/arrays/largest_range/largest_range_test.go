package main

import (
	"reflect"
	"runtime"
	"testing"
)

func TestLargestRange(t *testing.T) {

	funcs := []func([]int) [2]int{
		largestRange,
		largestRange1,
	}
	exp := [2]int{0, 7}

	for _, fn := range funcs {
		pc := reflect.ValueOf(fn).Pointer()
		fun := runtime.FuncForPC(pc)

		t.Run(fun.Name(), func(t *testing.T) {
			arr := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
			res := fn(arr)

			if !reflect.DeepEqual(res, exp) {
				t.Errorf("expected %v, but got %v", exp, res)
			}
		})
	}
}
