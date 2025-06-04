package main

import (
	"math/rand"
	"reflect"
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(50)
	}

	exp := make([]int, len(arr))
	copy(exp, arr)
	slices.Sort(exp)

	countingSort(arr)

	if !reflect.DeepEqual(exp, arr) {
		t.Errorf("expected %v, but got %v", exp, arr)
	}
}
