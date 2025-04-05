package main

import "testing"

func TestCountSubstrings(t *testing.T) {
	testTable := []struct {
		name   string
		input  string
		expRes int
	}{
		{
			name:   "happy path abc",
			input:  "abc",
			expRes: 3,
		},
		{
			name:   "happy path aaa",
			input:  "aaa",
			expRes: 6,
		},
		{
			name:   "happy path aabbbaa",
			input:  "aabbbaa",
			expRes: 14,
		},
		{
			name:   "happy path aaab",
			input:  "aaab",
			expRes: 7,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := countSubstrings(subtest.input)

			if subtest.expRes != res {
				t.Errorf("expected (%v), but got (%v)", subtest.expRes, res)
			}
		})
	}
}
