package main

import (
	"reflect"
	"runtime"
	"testing"
)

func TestValidateStartingCity(t *testing.T) {

	distances := []int{5, 25, 15, 10, 15}
	fuel := []int{1, 2, 1, 0, 3}
	mpg := 10
	expected := 4

	funcs := []func([]int, []int, int) int{
		validateStartingCity,
		validateStartingCity1,
	}

	for _, fn := range funcs {
		pc := reflect.ValueOf(fn).Pointer()
		fun := runtime.FuncForPC(pc)

		t.Run(fun.Name(), func(t *testing.T) {
			actual := fn(distances, fuel, mpg)

			if actual != expected {
				t.Errorf("expected %v, but got %v", expected, actual)
			}
		})
	}
}
