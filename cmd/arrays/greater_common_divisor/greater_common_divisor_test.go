package main

import "testing"

func TestGreaterCommonDivisor(t *testing.T) {
	testTable := []struct {
		name string
		x    int
		y    int
		exp  int
	}{
		{
			name: "gcd 18 and 6 = 6",
			x:    18,
			y:    6,
			exp:  6,
		},
		{
			name: "gcd 18 and 10 = 2",
			x:    18,
			y:    10,
			exp:  2,
		},
		{
			name: "gcd 17 and 11 = 1",
			x:    17,
			y:    11,
			exp:  1,
		},
		{
			name: "gcd 5 and 15 = 5",
			x:    5,
			y:    15,
			exp:  5,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := greaterCommonDivisor(subtest.x, subtest.y)

			if subtest.exp != res {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}
