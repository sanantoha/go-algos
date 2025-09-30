package main

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		exp  int
	}{
		{
			name: "subarraySum1",
			arr:  []int{1, 1, 1},
			k:    2,
			exp:  2,
		},
		{
			name: "subarraySum2",
			arr:  []int{1, 2, 3},
			k:    3,
			exp:  2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := subarraySum(test.arr, test.k); got != test.exp {
				t.Errorf("subarraySum() = %v, want %v", got, test.exp)
			}
		})
	}
}
