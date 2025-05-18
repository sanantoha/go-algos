package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	testTable := []struct {
		name  string
		input []rune
		exp   []rune
	}{
		{
			name:  "happy path",
			input: []rune{'h', 'e', 'l', 'l', 'o'},
			exp:   []rune{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:  "empty slice",
			input: []rune{},
			exp:   []rune{},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			reverseString(subtest.input)

			if !reflect.DeepEqual(subtest.exp, subtest.input) {
				t.Errorf("expected %v, but got %v", subtest.exp, subtest.input)
			}
		})
	}
}
