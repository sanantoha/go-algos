package main

import (
	"testing"
)

func TestMinNumberOfCoinsForChange(t *testing.T) {
	exp := 3
	input := []int{1, 5, 10}
	actual := minNumberOfCoinsForChange(7, input)

	if exp != actual {
		t.Errorf("expected %v, but got %v", exp, actual)
	}
}
