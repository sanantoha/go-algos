package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	testTable := []struct {
		name string
		n    int
		exp  int
	}{
		{
			name: "sqrt(4) == 2",
			n:    4,
			exp:  2,
		},
		{
			name: "sqrt(5) == 2",
			n:    5,
			exp:  2,
		},
		{
			name: "sqrt(6) == 2",
			n:    6,
			exp:  2,
		},
		{
			name: "sqrt(7) == 2",
			n:    7,
			exp:  2,
		},
		{
			name: "sqrt(8) == 2",
			n:    7,
			exp:  2,
		},
		{
			name: "sqrt(9) == 3",
			n:    9,
			exp:  3,
		},
		{
			name: "sqrt(16) == 4",
			n:    16,
			exp:  4,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := sqrt(subtest.n)

			if subtest.exp != res {
				t.Errorf("expected (%v), got (%v)", subtest.exp, res)
			}
		})
	}
}
