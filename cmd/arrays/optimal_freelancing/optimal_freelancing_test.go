package main

import (
	"testing"
)

func TestOptimalFreelancing(t *testing.T) {

	jobs := []struct {
		deadline int
		payment  int
	}{
		{
			deadline: 2,
			payment:  2,
		},
		{
			deadline: 4,
			payment:  3,
		},
		{
			deadline: 5,
			payment:  1,
		},
		{
			deadline: 7,
			payment:  2,
		},
		{
			deadline: 3,
			payment:  1,
		},
		{
			deadline: 3,
			payment:  2,
		},
		{
			deadline: 1,
			payment:  3,
		},
	}

	exp := 13

	res := optimalFreelancing(jobs)

	if res != exp {
		t.Errorf("expected: %d, got: %d", exp, res)
	}
}
