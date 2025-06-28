package main

import (
	"reflect"
	"testing"
)

func TestFindLadders(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	beginWord := "hit"
	endWord := "cog"

	exp := [][]string{
		{"hit", "hot", "dot", "dog", "cog"},
		{"hit", "hot", "lot", "log", "cog"},
	}

	res := findLadders(beginWord, endWord, wordList)

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("expected: %v, got: %v", exp, res)
	}
}
