package main

import "testing"

func TestWordBreak(t *testing.T) {
	testTable := []struct {
		name     string
		s        string
		wordDict []string
		exp      bool
	}{
		{
			name:     "word: leetcode",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			exp:      true,
		},
		{
			name:     "word: applepenapple",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			exp:      true,
		},
		{
			name:     "word: catsandog",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			exp:      false,
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			res := wordBreak(subtest.s, subtest.wordDict)

			if subtest.exp != res {
				t.Errorf("expected %v, got %v", subtest.exp, res)
			}
		})
	}
}
