package main

import (
	"testing"
)

func TestLadderLength(t *testing.T) {

	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		exp       int
	}{
		{
			name:      "happy path",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			exp:       5,
		},
		{
			name:      "happy path 1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			exp:       0,
		},
		{
			name:      "happy path 1",
			beginWord: "MAMA",
			endWord:   "SIRI",
			wordList:  []string{"SAMA", "SIMA", "SIRA", "SIRI", "MISA", "DISA"},
			exp:       5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ladderLength(tt.beginWord, tt.endWord, tt.wordList)

			if res != tt.exp {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
