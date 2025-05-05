package main

import (
	"reflect"
	"testing"
)

var arr = []int{1, 2, 3}
var exp = [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}

func TestPowerset(t *testing.T) {

	res := powerset(arr)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}

func TestPowersetRec(t *testing.T) {

	res := powersetRec(arr)

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("expected %v, but got %v", exp, res)
	}
}
