package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinimalHaviestSetA(t *testing.T) {
	exp := []int{5, 6}
	arr := []int{6, 4, 2, 3, 1, 5}

	res := minimalHaviestSetA(arr)
	fmt.Println(res)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
