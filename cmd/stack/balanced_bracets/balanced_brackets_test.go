package main

import "testing"

func TestBalancedBrackets(t *testing.T) {
	testTable := []struct {
		name  string
		input string
		exp   bool
	}{
		{
			name:  "happy path",
			input: "([])(){}(())()()",
			exp:   true,
		},
		{
			name:  "wrong bracket",
			input: "}{",
			exp:   false,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := balancedBrackets(subtest.input)
			if subtest.exp != res {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}
