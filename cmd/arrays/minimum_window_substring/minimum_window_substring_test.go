package main

import "testing"

func TestMinimumWindowSubstring(t *testing.T) {
	testTable := []struct {
		name string
		s    string
		t    string
		exp  string
	}{
		{
			name: "happy path",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			exp:  "BANC",
		},
		{
			name: "happy path1",
			s:    "a",
			t:    "a",
			exp:  "a",
		},
		{
			name: "happy path1",
			s:    "a",
			t:    "aa",
			exp:  "",
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := minimumWindowSubstring(subtest.s, subtest.t)

			if res != subtest.exp {
				t.Errorf("expected %v, but got %v", subtest.exp, res)
			}
		})
	}
}
