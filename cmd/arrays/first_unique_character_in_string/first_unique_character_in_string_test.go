package main

import "testing"

func TestFirstUniqChar(t *testing.T) {
	testTable := []struct {
		name   string
		input  string
		expRes int
	}{
		{
			name:   "happy path leetcode",
			input:  "leetcode",
			expRes: 0,
		},
		{
			name:   "happy path loveleetcode",
			input:  "loveleetcode",
			expRes: 2,
		},
		{
			name:   "happy path aabb",
			input:  "aabb",
			expRes: -1,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := firstUniqChar(subtest.input)

			if subtest.expRes != res {
				t.Errorf("expected (%v), got (%v)", subtest.expRes, res)
			}
		})
	}
}
