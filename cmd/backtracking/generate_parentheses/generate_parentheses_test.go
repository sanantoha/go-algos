package main

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	tests := []struct {
		name string
		cnt  int
		exp  []string
	}{
		{
			name: "happy path cnt = 3",
			cnt:  3,
			exp:  []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "happy path cnt = 2",
			cnt:  2,
			exp:  []string{"(())", "()()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := generateParentheses(tt.cnt)

			if !reflect.DeepEqual(tt.exp, res) {
				t.Errorf("expected %v, but got %v", tt.exp, res)
			}
		})
	}
}
