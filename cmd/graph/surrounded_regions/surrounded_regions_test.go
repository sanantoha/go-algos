package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	testTable := []struct {
		name  string
		input [][]rune
		exp   [][]rune
	}{
		{
			name: "happy path",
			input: [][]rune{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			exp: [][]rune{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "happy path 1",
			input: [][]rune{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			exp: [][]rune{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "happy path 2",
			input: [][]rune{
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
			},
			exp: [][]rune{
				{'X', 'O', 'X', 'O', 'X', 'O'},
				{'O', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'X'},
			},
		},
	}

	for _, subtest := range testTable {
		t.Run(subtest.name, func(t *testing.T) {
			solve(subtest.input)

			if !reflect.DeepEqual(subtest.exp, subtest.input) {
				t.Errorf("expected\n%v	but got\n%v", runeMatrixToString(subtest.exp), runeMatrixToString(subtest.input))
			}
		})
	}
}

func runeMatrixToString(matrix [][]rune) string {
	var builder strings.Builder
	for _, row := range matrix {
		for _, r := range row {
			builder.WriteRune(r)
			builder.WriteRune(' ')
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}
