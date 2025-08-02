package main

import (
	"reflect"
	"runtime"
	"testing"
)

func TestSameBst(t *testing.T) {

	funcs := []func([]int, []int) bool{
		sameBsts,
		sameBsts1,
	}

	arr1 := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
	arr2 := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}

	for _, fn := range funcs {
		pc := reflect.ValueOf(fn).Pointer()
		fun := runtime.FuncForPC(pc)

		t.Run(fun.Name(), func(t *testing.T) {
			res := fn(arr1, arr2)

			if !res {
				t.Errorf("expected true, but got %v", res)
			}
		})
	}
}
